package status

import "github.com/whoisnian/glb/httpd"

func ReadyzHandler(store *httpd.Store) {
	store.Respond200(nil)
}

func LivezHandler(store *httpd.Store) {
	store.Respond200(nil)
}
