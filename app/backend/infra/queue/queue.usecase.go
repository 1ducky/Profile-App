package queue

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Queue[Key comparable, Jobs any, Reply any] struct {
	identity string
	pending  map[Key]*JobChan[Jobs, Reply]
	Job      chan *JobChan[Jobs, Reply]
	queue    chan Key
	closed   bool
	mu       *sync.RWMutex
	timeout  time.Duration
}

func NewQueue[Key comparable, Jobs any, Reply any](identity string, size int, timeoutSecond int) *Queue[Key, Jobs, Reply] {
	if timeoutSecond == 0 {
		timeoutSecond = 1
	}
	return &Queue[Key, Jobs, Reply]{
		identity: identity,
		pending:  make(map[Key]*JobChan[Jobs, Reply]),
		Job:      make(chan *JobChan[Jobs, Reply]),
		queue:    make(chan Key, size),
		closed:   false,
		mu:       &sync.RWMutex{},
		timeout:  time.Duration(timeoutSecond) * time.Second,
	}
}

func (q *Queue[Key, Jobs, Reply]) Start(ctx context.Context) {
	go func() {
		for pen := range q.queue {
			q.claimJob(ctx, pen)
		}
	}()
}

func (q *Queue[Key, Jobs, Reply]) DeletePending(key Key, expected *JobChan[Jobs, Reply]) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	current, exists := q.pending[key]
	if !exists || current != expected {
		return nil
	}
	close(q.pending[key].Reply)
	delete(q.pending, key)
	return nil
}

func (q *Queue[Key, Jobs, Reply]) claimJob(ctx context.Context, key Key) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, q.timeout)
	defer cancel()

	q.mu.Lock()
	pen, exits := q.pending[key]
	if !exits {
		q.mu.Unlock()
		return
	}
	delete(q.pending, key)
	q.mu.Unlock()
	select {
	case q.Job <- pen:
		return
	case <-ctxWithTimeout.Done():
		return
	}

}

func (q *Queue[Key, Jobs, Reply]) Close() error {
	q.mu.Lock()
	defer q.mu.Unlock()
	close(q.queue)
	close(q.Job)
	q.closed = true
	return nil
}

func (q *Queue[Key, Jobs, Reply]) Push(ctx context.Context, key Key, value Jobs) (Reply, error) {
	ctx, cancel := context.WithTimeout(ctx, q.timeout)
	defer cancel()
	res, err := q.Enqueue(ctx, key, value)
	if err != nil {
		return *new(Reply), err
	}

	select {
	case reply, ok := <-res:
		if !ok {
			return *new(Reply), fmt.Errorf("job cancelled")
		}
		return reply, nil

	case <-ctx.Done():
		return *new(Reply), fmt.Errorf("request Time out")
	}
}

func (q *Queue[Key, Jobs, Reply]) Enqueue(ctx context.Context, key Key, value Jobs) (<-chan Reply, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, q.timeout)
	defer cancel()

	q.mu.Lock()

	if q.closed {
		q.mu.Unlock()
		res := make(chan Reply)
		close(res)
		return res, fmt.Errorf("%s: Service Unavaliable", q.identity)
	}

	current, exits := q.pending[key]
	if exits {
		current.Job = value
		q.mu.Unlock()
		return current.Reply, nil
	}
	pending := &JobChan[Jobs, Reply]{Job: value, Reply: make(chan Reply)}
	q.pending[key] = pending
	q.mu.Unlock()

	// _, exists := q.pending[key]
	// pending := &JobChan[Jobs, Reply]{Job: value, Reply: make(chan Reply)}

	// // Always keep latest value.
	// q.pending[key] = pending

	// q.mu.Unlock()

	// if exists {
	// 	// Already has a slot in FIFO.
	// 	res := make(chan Reply)
	// 	close(res)
	// 	return res, nil
	// }

	select {
	case q.queue <- key:
		return pending.Reply, nil

	case <-ctxWithTimeout.Done():
		res := make(chan Reply)
		close(res)
		return res, fmt.Errorf("Timeout")
	}

}
