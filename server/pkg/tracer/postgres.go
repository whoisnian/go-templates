package tracer

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/whoisnian/glb/logger"
)

type PgTracer struct {
	LOG *logger.Logger
}

type pgTraceCtxKey int

const (
	_ pgTraceCtxKey = iota
	pgTraceQueryCtxKey
	pgTraceBatchCtxKey
	pgTraceCopyFromCtxKey
)

type traceQueryData struct {
	start time.Time
	sql   string
	args  []any
}

func (*PgTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	return context.WithValue(ctx, pgTraceQueryCtxKey, &traceQueryData{
		start: time.Now(),
		sql:   data.SQL,
		args:  data.Args,
	})
}

func (pt *PgTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	queryData := ctx.Value(pgTraceQueryCtxKey).(*traceQueryData)
	duration := time.Since(queryData.start)

	result := "OK"
	if data.Err != nil {
		result = data.Err.Error()
	}

	pt.LOG.Debug(ctx, "PG.TraceQueryEnd",
		slog.Attr{
			Key: "pg",
			Value: slog.GroupValue(
				slog.Duration("dur", duration),
				slog.String("result", result),
				slog.String("sql", queryData.sql),
				slog.Any("args", queryData.args),
			),
		},
	)
}

type traceBatchData struct {
	start time.Time
	batch *pgx.Batch
}

func (*PgTracer) TraceBatchStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchStartData) context.Context {
	return context.WithValue(ctx, pgTraceBatchCtxKey, &traceBatchData{
		start: time.Now(),
		batch: data.Batch,
	})
}

func (pt *PgTracer) TraceBatchQuery(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchQueryData) {
	result := "OK"
	if data.Err != nil {
		result = data.Err.Error()
	}

	pt.LOG.Debug(ctx, "PG.TraceBatchQuery",
		slog.Attr{
			Key: "pg",
			Value: slog.GroupValue(
				slog.String("result", result),
				slog.String("sql", data.SQL),
				slog.Any("args", data.Args),
			),
		},
	)
}

func (pt *PgTracer) TraceBatchEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchEndData) {
	batchData := ctx.Value(pgTraceBatchCtxKey).(*traceBatchData)
	duration := time.Since(batchData.start)

	result := "OK"
	if data.Err != nil {
		result = data.Err.Error()
	}

	pt.LOG.Debug(ctx, "PG.TraceBatchEnd",
		slog.Attr{
			Key: "pg",
			Value: slog.GroupValue(
				slog.Duration("dur", duration),
				slog.String("result", result),
				slog.Int("size", batchData.batch.Len()),
			),
		},
	)
}

type traceCopyFromData struct {
	start       time.Time
	tableName   pgx.Identifier
	columnNames []string
}

func (*PgTracer) TraceCopyFromStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceCopyFromStartData) context.Context {
	return context.WithValue(ctx, pgTraceCopyFromCtxKey, &traceCopyFromData{
		start:       time.Now(),
		tableName:   data.TableName,
		columnNames: data.ColumnNames,
	})
}

func (pt *PgTracer) TraceCopyFromEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceCopyFromEndData) {
	copyFromData := ctx.Value(pgTraceCopyFromCtxKey).(*traceCopyFromData)
	duration := time.Since(copyFromData.start)

	result := "OK"
	if data.Err != nil {
		result = data.Err.Error()
	}

	pt.LOG.Debug(ctx, "PG.TraceCopyFromEnd",
		slog.Attr{
			Key: "pg",
			Value: slog.GroupValue(
				slog.Duration("dur", duration),
				slog.String("result", result),
				slog.String("table", copyFromData.tableName.Sanitize()),
				slog.Any("columns", copyFromData.columnNames),
			),
		},
	)
}
