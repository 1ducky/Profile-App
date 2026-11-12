package contextapp

import (
	"context"
)

func GetTraceID(ctx context.Context) string {
	if val, ok := ctx.Value(TraceIDKey).(string); ok {
		return val
	}
	return ""
}

func WithTracer(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDKey, traceID)
}

func WithAuth(ctx context.Context, auth SessionContext) context.Context {
	return context.WithValue(ctx, AuthUserKey, auth)
}

func GetAuth(ctx context.Context) (SessionContext, bool) {
	user, ok := ctx.Value(AuthUserKey).(SessionContext)
	return user, ok
}
