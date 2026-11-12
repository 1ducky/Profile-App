package queue

type QueueStatus string

const (
	Closed  QueueStatus = "Closed"
	Already QueueStatus = "Already"
	Ready   QueueStatus = "Ready"
)
