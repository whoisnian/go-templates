package tracer

import (
	"context"
	"log/slog"
)

type WrapHandler struct{ slog.Handler }

func (h *WrapHandler) Handle(ctx context.Context, r slog.Record) error {
	traceID := "unknown"
	spanID := "unknown"
	if sc := SpanContextFromContext(ctx); sc.IsValid() {
		traceID = sc.TraceID().String()
		spanID = sc.SpanID().String()
	}
	r.AddAttrs(
		slog.String("trace_id", traceID),
		slog.String("span_id", spanID),
	)
	return h.Handler.Handle(ctx, r)
}
