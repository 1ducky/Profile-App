package transport

import (
	"encoding/json"
	"log"
	"net/http"
	"profile-template/internal/app/err"
	"profile-template/internal/contextapp"
)

type Adapter struct {
	errApp *err.ErrApp
}

func NewAdapter(errApp *err.ErrApp) *Adapter {
	return &Adapter{
		errApp: errApp,
	}
}

func (a *Adapter) Adapt(f HandlerFn, domain err.ErrDomain) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if errReq := f(w, r); errReq != nil {
			entry := a.errApp.Translate(errReq, domain)
			if entry.Code == err.InternalErrCode {
				log.Print("Invariant Detected : ", errReq, " requestid: ", contextapp.GetTraceID(r.Context()))
			}
			if entry.StatusCode == 0 {
				entry.StatusCode = http.StatusOK
			}
			w.WriteHeader(entry.StatusCode)
			json.NewEncoder(w).Encode(ErrorResponse{
				Code:    string(entry.Code),
				Message: entry.Message,
				TraceID: contextapp.GetTraceID(r.Context()),
				Status:  StatusError,
			})
		}

	}
}
