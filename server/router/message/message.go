package message

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/whoisnian/glb/httpd"
	"github.com/whoisnian/glb/logger"
	"github.com/whoisnian/go-templates/server/global"
	"github.com/whoisnian/go-templates/server/model"
)

func ListHandler(store *httpd.Store) {
	const sql = `SELECT id, content, created_at FROM messages`
	rows, err := global.DB.Query(store.R.Context(), sql)
	if err != nil {
		global.LOG.Error(store.R.Context(), "pgxpool.Query", logger.Error(err))
		store.W.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []model.Message
	for rows.Next() {
		var msg model.Message
		if err = rows.Scan(&msg.Id, &msg.Content, &msg.CreatedAt); err != nil {
			global.LOG.Error(store.R.Context(), "rows.Scan", logger.Error(err))
			store.W.WriteHeader(http.StatusInternalServerError)
			return
		}
		results = append(results, msg)
	}
	if err = rows.Err(); err != nil {
		global.LOG.Error(store.R.Context(), "rows.Err", logger.Error(err))
		store.W.WriteHeader(http.StatusInternalServerError)
		return
	}
	store.RespondJson(results)
}

func CreateHandler(store *httpd.Store) {
	var msg model.Message
	err := json.NewDecoder(store.R.Body).Decode(&msg)
	if err != nil {
		store.W.WriteHeader(http.StatusBadRequest)
		store.W.Write([]byte(err.Error()))
		return
	}

	const sql = `INSERT INTO messages (content) VALUES ($1) RETURNING id, content, created_at`
	row := global.DB.QueryRow(store.R.Context(), sql, msg.Content)
	if err = row.Scan(&msg.Id, &msg.Content, &msg.CreatedAt); err != nil {
		global.LOG.Error(store.R.Context(), "rows.Scan", logger.Error(err))
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

	const sql = `DELETE FROM messages WHERE id = $1`
	if _, err = global.DB.Exec(store.R.Context(), sql, id); err != nil {
		global.LOG.Error(store.R.Context(), "pgxpool.Exec", logger.Error(err))
		store.W.WriteHeader(http.StatusInternalServerError)
		return
	}
	store.RespondJson(nil)
}
