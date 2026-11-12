package contextapp

type contextKey string

const (
	TraceIDKey  contextKey = "traceID"
	AuthUserKey contextKey = "auth"
)
