package web

import (
	"context"
)

type ctxKey int

const traceID ctxKey = 1

func setTraceID(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, traceID, value)
}

func GetTraceID(ctx context.Context) string {
	v, ok := ctx.Value(traceID).(string)
	if !ok {
		return "00000000-0000-0000-0000-000000000000"
	}

	return v
}
