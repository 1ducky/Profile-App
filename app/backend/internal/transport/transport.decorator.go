package transport

import "net/http"

type TrasnportWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (rw *TrasnportWriter) WriteHeader(code int) {
	rw.StatusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
