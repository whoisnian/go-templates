package tracer

import "github.com/whoisnian/glb/httpd"

func NewHttpdMiddleware() func(*httpd.Store) {
	return func(store *httpd.Store) {
		store.Next()
	}
}
