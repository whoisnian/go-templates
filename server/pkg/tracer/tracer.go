package tracer

import (
	"context"
)

type SpanContext struct {
	ctx context.Context
}

func SpanContextFromContext(ctx context.Context) SpanContext {
	return SpanContext{ctx: ctx}
}
