package tracer

import (
	"github.com/whoisnian/glb/httpd"
)

func NewHttpdMiddleware() func(*httpd.Store) {
	return func(store *httpd.Store) {
		trace, ctx := StartTraceFromContext(store.R.Context(), store.I.Method+" "+store.I.Path)
		defer trace.End()

		store.R = store.R.WithContext(ctx)
		store.Next()
	}
}
