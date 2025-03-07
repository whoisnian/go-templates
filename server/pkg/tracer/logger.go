package tracer

import (
	"context"
	"log/slog"
)

type WrapHandler struct{ slog.Handler }

func (h *WrapHandler) Handle(ctx context.Context, r slog.Record) error {
	tid := TraceFromContext(ctx).ID()
	r.AddAttrs(
		slog.String("trace_id", tid),
	)
	return h.Handler.Handle(ctx, r)
}
