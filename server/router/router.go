package router

import (
	"context"
	"net/http"

	"github.com/whoisnian/glb/httpd"
	"github.com/whoisnian/go-templates/server/global"
	"github.com/whoisnian/go-templates/server/router/message"
	"github.com/whoisnian/go-templates/server/router/status"
)

func Setup(_ context.Context) *httpd.Mux {
	mux := httpd.NewMux()
	mux.HandleMiddleware(global.LOG.NewMiddleware())

	mux.Handle("/api/messages", http.MethodGet, message.ListHandler)
	mux.Handle("/api/messages", http.MethodPost, message.CreateHandler)
	mux.Handle("/api/messages/:id", http.MethodDelete, message.DeleteHandler)

	mux.Handle("/readyz", http.MethodGet, status.ReadyzHandler)
	mux.Handle("/livez", http.MethodGet, status.LivezHandler)
	return mux
}
