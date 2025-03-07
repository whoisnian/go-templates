package tracer

import (
	"context"
	crand "crypto/rand"
	"encoding/binary"
	"math/rand/v2"
	"strconv"
	"sync"
	"time"
)

type Trace struct {
	id      uint64
	name    string
	startAt time.Time
	endAt   time.Time
}

var emptyTrace = &Trace{}

func (t *Trace) IsEmpty() bool { return t.id == 0 }
func (t *Trace) ID() string    { return strconv.FormatUint(t.id, 16) }
func (t *Trace) Name() string  { return t.name }
func (t *Trace) End()          { t.endAt = time.Now() }

type traceContextKeyType struct{}

var traceContextKey = traceContextKeyType{}

func StartTraceFromContext(ctx context.Context, name string) (*Trace, context.Context) {
	if ctx == nil {
		ctx = context.Background()
	} else if trace, ok := ctx.Value(traceContextKey).(*Trace); ok {
		return trace, ctx
	}

	trace := &Trace{
		id:      generateTraceID(),
		name:    name,
		startAt: time.Now(),
	}
	return trace, context.WithValue(ctx, traceContextKey, trace)
}

func TraceFromContext(ctx context.Context) *Trace {
	if ctx == nil {
		return emptyTrace
	}
	if trace, ok := ctx.Value(traceContextKey).(*Trace); ok {
		return trace
	}
	return emptyTrace
}

var tidMutex sync.Mutex
var tidSource = rand.NewPCG(func() (seed1, seed2 uint64) {
	_ = binary.Read(crand.Reader, binary.LittleEndian, &seed1)
	_ = binary.Read(crand.Reader, binary.LittleEndian, &seed2)
	return
}())

func generateTraceID() uint64 {
	tidMutex.Lock()
	defer tidMutex.Unlock()
	return tidSource.Uint64()
}
