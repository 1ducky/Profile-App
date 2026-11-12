package pipeline

import (
	"context"
	"sync"
)

type WorkerPool[J any, R any] struct {
	poolSize int
}

func NewWorkerPool[J any, R any](poolSize int) *WorkerPool[J, R] {
	return &WorkerPool[J, R]{poolSize: poolSize}
}

func (w *WorkerPool[J, R]) Run(ctx context.Context, Job <-chan J, fn func(context.Context, J) R) <-chan R {
	var wg sync.WaitGroup

	ResultStream := make(chan R)
	for workerID := range w.poolSize {
		wg.Add(1)
		go func(ID int) {

			defer wg.Done()

		loop:
			for {
				select {
				case <-ctx.Done():
					break loop
				case jobItem, ok := <-Job:
					if !ok {
						break loop
					}

					result := fn(ctx, jobItem)
					ResultStream <- result
				}

			}

		}(workerID)
	}

	go func() {
		wg.Wait()
		close(ResultStream)
	}()
	return ResultStream

}
