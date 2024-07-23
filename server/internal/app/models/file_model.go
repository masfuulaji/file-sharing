package models

import (
	"time"

	"github.com/lib/pq"
)

type File struct {
	Id        string      `json:"id"`
	FileName  string      `db:"file_name" json:"file_name"`
	FilePath  string      `db:"file_path" json:"file_path"`
	UserId    string      `db:"user_id" json:"user_id"`
	CreatedAt time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt time.Time   `db:"updated_at" json:"updated_at"`
	DeletedAt pq.NullTime `db:"deleted_at" json:"deleted_at"`
}
