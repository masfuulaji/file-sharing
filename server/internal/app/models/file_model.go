package models

import (
	"time"

	"github.com/lib/pq"
)

type File struct {
	Id        string      `json:"id"`
	FileName  string      `json:"file_name"`
	UserId    string      `json:"user_id"`
	CreatedAt time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt time.Time   `db:"updated_at" json:"updated_at"`
	DeletedAt pq.NullTime `db:"deleted_at" json:"deleted_at"`
}
