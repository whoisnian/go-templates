package postgres

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/whoisnian/glb/logger"
)

type Tracer struct {
	LOG *logger.Logger
}

type ctxKey int

const (
	_ ctxKey = iota
	tracerQueryCtxKey
	tracerBatchCtxKey
	tracerCopyFromCtxKey
)

type tracerQueryData struct {
	start time.Time
	sql   string
	args  []any
}

func (*Tracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	return context.WithValue(ctx, tracerQueryCtxKey, &tracerQueryData{
		start: time.Now(),
		sql:   data.SQL,
		args:  data.Args,
	})
}

func (t *Tracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	queryData := ctx.Value(tracerQueryCtxKey).(*tracerQueryData)
	interval := time.Since(queryData.start)

	result := "OK"
	if data.Err != nil {
		result = data.Err.Error()
	}

	t.LOG.Debug("PG Query",
		slog.Duration("interval", interval),
		slog.String("result", result),
		slog.String("sql", queryData.sql),
		slog.Any("args", queryData.args),
	)
}

type tracerBatchData struct {
	start time.Time
	batch *pgx.Batch
}

func (*Tracer) TraceBatchStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchStartData) context.Context {
	return context.WithValue(ctx, tracerBatchCtxKey, &tracerBatchData{
		start: time.Now(),
		batch: data.Batch,
	})
}

func (t *Tracer) TraceBatchQuery(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchQueryData) {
	result := "OK"
	if data.Err != nil {
		result = data.Err.Error()
	}

	t.LOG.Debug("PG BatchQuery",
		slog.String("result", result),
		slog.String("sql", data.SQL),
		slog.Any("args", data.Args),
	)
}

func (t *Tracer) TraceBatchEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchEndData) {
	batchData := ctx.Value(tracerBatchCtxKey).(*tracerBatchData)
	interval := time.Since(batchData.start)

	result := "OK"
	if data.Err != nil {
		result = data.Err.Error()
	}

	t.LOG.Debug("PG Batch",
		slog.Duration("interval", interval),
		slog.String("result", result),
		slog.Int("size", batchData.batch.Len()),
	)
}

type tracerCopyFromData struct {
	start       time.Time
	tableName   pgx.Identifier
	columnNames []string
}

func (*Tracer) TraceCopyFromStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceCopyFromStartData) context.Context {
	return context.WithValue(ctx, tracerCopyFromCtxKey, &tracerCopyFromData{
		start:       time.Now(),
		tableName:   data.TableName,
		columnNames: data.ColumnNames,
	})
}

func (t *Tracer) TraceCopyFromEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceCopyFromEndData) {
	copyFromData := ctx.Value(tracerCopyFromCtxKey).(*tracerCopyFromData)
	interval := time.Since(copyFromData.start)

	result := "OK"
	if data.Err != nil {
		result = data.Err.Error()
	}

	t.LOG.Debug("PG CopyFrom",
		slog.Duration("interval", interval),
		slog.String("result", result),
		slog.String("table", copyFromData.tableName.Sanitize()),
		slog.Any("columns", copyFromData.columnNames),
	)
}
