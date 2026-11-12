package transport

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"profile-template/internal/config"
)

func NewTransport(cfg config.HTTPConfig, handler http.Handler) TransportAPI {
	log.Print(cfg)

	return &Transport{

		server: &http.Server{
			Addr:              net.JoinHostPort(cfg.Host, cfg.Port), // Address to listen on
			Handler:           handler,                              // Handler for HTTP requests
			ReadTimeout:       cfg.ReadTimeout,                      // Maximum duration for reading the entire request
			ReadHeaderTimeout: cfg.ReadHeaderTimeout,                // Maximum duration for reading the request headers
			WriteTimeout:      cfg.WriteTimeout,                     // Maximum duration for writing the response
			IdleTimeout:       cfg.IdleTimeout,                      // Maximum duration for keeping the connection open
		},
		Handler: handler,
	}

}

// ListenAndServe HTTP Server
// It will run in background (goroutine)
func (t *Transport) Start() error {
	err := t.server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

// Shutdown HTTP Server
// With Timeout Context to wait for the server to finish all request before close
func (t *Transport) Stop(ctx context.Context) error {
	return t.server.Shutdown(ctx)
}

func (t *Transport) Address() string {
	return t.server.Addr
}
