package sse

import (
	"context"
)

type Connection struct {
	Reply  chan []byte
	Ctx    context.Context
	cancel context.CancelFunc
}

func NewConnection(ctx context.Context) *Connection {
	newCtx, cancel := context.WithCancel(ctx)
	return &Connection{
		Reply:  make(chan []byte, 100),
		Ctx:    newCtx,
		cancel: cancel,
	}
}

func (c *Connection) Send(message []byte) error {
	select {
	case <-c.Ctx.Done():
		return ErrConnClosed
	case c.Reply <- message:
		return nil
	default:
		return ErrNoMessageSend
	}
}

func (c *Connection) Close() {
	c.cancel()
}
