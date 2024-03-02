package message

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/whoisnian/glb/httpd"
	"github.com/whoisnian/go-templates/cli/global"
	"github.com/whoisnian/go-templates/server/model/message"
)

func ListHandler(store *httpd.Store) {
	msgs, err := message.List(store.R.Context())
	if err != nil {
		global.LOG.Error(err.Error())
		store.W.WriteHeader(http.StatusInternalServerError)
		return
	}
	store.RespondJson(msgs)
}

func CreateHandler(store *httpd.Store) {
	var msg message.Message
	err := json.NewDecoder(store.R.Body).Decode(&msg)
	if err != nil {
		store.W.WriteHeader(http.StatusBadRequest)
		store.W.Write([]byte(err.Error()))
		return
	}
	msg, err = message.Create(store.R.Context(), msg.Content)
	if err != nil {
		global.LOG.Error(err.Error())
		store.W.WriteHeader(http.StatusInternalServerError)
		return
	}
	store.RespondJson(msg)
}

func DeleteHandler(store *httpd.Store) {
	id, err := strconv.ParseInt(store.RouteParam("id"), 10, 64)
	if err != nil {
		store.W.WriteHeader(http.StatusBadRequest)
		store.W.Write([]byte("invalid ID"))
		return
	}
	if err = message.DeleteById(store.R.Context(), id); err != nil {
		global.LOG.Error(err.Error())
		store.W.WriteHeader(http.StatusInternalServerError)
		return
	}
	store.RespondJson(nil)
}
