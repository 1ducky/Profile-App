package sse

import "errors"

var (
	ErrNoMessageSend = errors.New("No Message Send")
	ErrConnClosed    = errors.New("Connection Closed")
	ErrConnNotFound  = errors.New("Connection Not Found")
	ErrTimeout       = errors.New("Timeout")
)
