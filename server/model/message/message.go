package message

import (
	"context"
	"time"

	"github.com/whoisnian/go-templates/server/global"
)

type Message struct {
	Id        int64     `db:"id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

const sqlCreate = `INSERT INTO messages (content) VALUES ($1) RETURNING id, content, created_at`

func Create(ctx context.Context, content string) (m Message, err error) {
	row := global.DB.QueryRow(ctx, sqlCreate, content)
	err = row.Scan(&m.Id, &m.Content, &m.CreatedAt)
	return m, err
}

const sqlDeleteById = `DELETE FROM messages WHERE id = $1`

func DeleteById(ctx context.Context, id int64) (err error) {
	_, err = global.DB.Exec(ctx, sqlDeleteById, id)
	return err
}

const sqlList = `SELECT id, content, created_at FROM messages`

func List(ctx context.Context) (result []Message, err error) {
	rows, err := global.DB.Query(ctx, sqlList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m Message
		if err = rows.Scan(&m.Id, &m.Content, &m.CreatedAt); err != nil {
			return result, err
		}
		result = append(result, m)
	}
	return result, rows.Err()
}
