package middleware

import (
	"log"
	"net/http"
	"profile-template/internal/contextapp"
	"profile-template/internal/transport"
	"time"

	"github.com/google/uuid"
)

func TrackerMiddleware() func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			// TODO : genorate request id and add to context
			requestId := uuid.New().String()
			ctxtracer := contextapp.WithTracer(r.Context(), requestId)
			r = r.WithContext(ctxtracer)
			tw := &transport.TrasnportWriter{ResponseWriter: w}

			// TODO : add request id to response header
			w.Header().Set("x-request-id", requestId)
			// TODO : add request id to logger
			next.ServeHTTP(tw, r)
			endReq := time.Since(startTime)

			log.Print("requestid: ", requestId, "", " [", tw.StatusCode, "]  path: ", r.URL.Path, " method: ", r.Method, " cost: ", endReq)
		})
	}
}
