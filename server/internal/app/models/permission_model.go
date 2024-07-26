package models

import (
	"time"

	"github.com/lib/pq"
)

type Permission struct {
	Id         string      `json:"id"`
	UserId     string      `db:"user_id" json:"user_id"`
	FileId     string      `db:"file_id" json:"file_id"`
	Permission string      `json:"permission"`
	CreatedAt  time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time   `db:"updated_at" json:"updated_at"`
	DeletedAt  pq.NullTime `db:"deleted_at" json:"deleted_at"`
}
