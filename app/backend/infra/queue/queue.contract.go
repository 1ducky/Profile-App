package queue

type JobChan[J any, V any] struct {
	Reply    chan V
	Job      J
	Proceess func(J) (V, error)
}

type WorkerStatus struct {
	id     int
	isDead bool
	caused error
}
