package sse

import (
	"context"
	"profile-template/internal/utils"
	"sync"
)

type Usecase struct {
	Conn map[string][]*Connection
	mu   sync.RWMutex
}

func NewUsecase() *Usecase {
	return &Usecase{
		Conn: make(map[string][]*Connection),
	}
}

func (u *Usecase) Connect(ctx context.Context, identifyer string) *Connection {
	conn := NewConnection(ctx)
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Conn[identifyer] = append(u.Conn[identifyer], conn)
	return conn
}

func (u *Usecase) Disconnect(ctx context.Context, identifyer string, conn *Connection) {
	u.mu.Lock()
	defer u.mu.Unlock()
	conn.Close()
	res := utils.MapField(u.Conn[identifyer], func(item *Connection) (*Connection, bool) {
		if item == conn {
			return nil, false // Remove if Address are Match
		}
		return item, true // Keep if Address are not Match
	})
	if len(res) == 0 {
		delete(u.Conn, identifyer)
		return
	}
	u.Conn[identifyer] = res
}

func (u *Usecase) SendMessage(ctx context.Context, identifyer string, message []byte) error {
	u.mu.RLock()
	conn, exits := u.Conn[identifyer]
	if !exits {
		u.mu.RUnlock()
		return ErrConnNotFound
	}
	snapshoot := make([]*Connection, 0, len(conn))
	snapshoot = append(snapshoot, conn...)
	u.mu.RUnlock()
	for _, c := range conn {
		c.Send(message)
	}

	return nil
}
