package transport

import (
	"context"
	"net/http"
)

type TransportAPI interface {
	Start() error
	Stop(context.Context) error
	Address() string
}

type Transport struct {
	server  *http.Server
	Handler http.Handler
}

type HandlerFn func(w http.ResponseWriter, r *http.Request) error

type Status string

const (
	StatusSuccess Status = "success"
	StatusError   Status = "error"
)

type Response[T any] struct {
	Status Status `json:"status"`
	Data   T      `json:"data"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	TraceID string `json:"request_id"`
	Status  Status `json:"status"`
}
