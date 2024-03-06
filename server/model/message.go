package model

import "time"

type Message struct {
	Id        int64     `db:"id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}
