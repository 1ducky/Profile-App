package sse_test

import (
	"context"
	"log"
	"sync"
	"testing"

	sse "profile-template/infra/SSE"
)

func TestConnect(t *testing.T) {

	manager := sse.NewUsecase()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// var wg sync.WaitGroup
	ConnCount := 10
	var conns []*sse.Connection
	user := []string{"User1", "User2", "User3"}

	for _, u := range user {
		for i := 0; i < ConnCount; i++ {
			conns = append(conns, manager.Connect(ctx, u))
		}
	}
	if len(conns) != (len(user) * ConnCount) {
		t.Errorf("Expected %d connections, got %d", len(user)*ConnCount, len(conns))
	}
	log.Printf("Koneksi Total %d", len(conns))

	for _, u := range user {
		for _, c := range conns {
			log.Println("Close Connection : ", u, c)
			manager.Disconnect(ctx, u, c)
		}
	}

}

func TestSendMSG(t *testing.T) {
	manager := sse.NewUsecase()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	var mu sync.RWMutex

	var deadWorker []string
	ConnCount := 50
	var conns []*sse.Connection
	user := []string{"User1", "User2", "User3", "user4", "user5", "user6"}

	for _, u := range user {
		for i := 0; i < ConnCount; i++ {
			conn := manager.Connect(ctx, u)
			conns = append(conns, conn)
			wg.Add(1)
			go func(conn *sse.Connection, u string) {
				defer wg.Done()
				for {
					select {
					case msg := <-conn.Reply:
						log.Println("Received message : ", string(msg), "User : ", u)
					case <-conn.Ctx.Done():
						mu.Lock()
						deadWorker = append(deadWorker, u)
						mu.Unlock()
						return
					}
				}
			}(conn, u)
		}
	}
	if len(conns) != (len(user) * ConnCount) {
		t.Errorf("Expected %d connections, got %d", len(user)*ConnCount, len(conns))
	}
	log.Printf("Koneksi Total %d", len(conns))

	// Send message
	for _, u := range user {
		go func(usr string) {
			err := manager.SendMessage(ctx, usr, []byte("Hello"))
			if err != nil {
				t.Errorf("Error sending message: %v", err)
			}
		}(u)
		go func(usr string) {
			err := manager.SendMessage(ctx, usr, []byte("World"))
			if err != nil {
				t.Errorf("Error sending message: %v", err)
			}
		}(u)
	}
	for _, u := range user {
		for _, c := range conns {
			go func() {

				manager.Disconnect(ctx, u, c)
			}()
		}
	}

	wg.Wait()
	mu.RLock()
	if len(deadWorker) != (len(user) * ConnCount) {
		t.Errorf("Expected %d dead workers, got %d", len(user)*ConnCount, len(deadWorker))
	}
	log.Print(len(deadWorker))
	err := manager.SendMessage(ctx, "user1", []byte("hello"))
	if err == nil {
		t.Errorf("should be error")
	}
	log.Print("Error Sending : ", err)
	mu.RUnlock()

}
