package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/file-sharing/internal/app/models"
)

var file models.File

type FileRepository interface {
	CreateFile(file models.File) error
	GetFile(id string) (models.File, error)
	GetFiles() ([]models.File, error)
	UpdateFile(file models.File, id string) error
	DeleteFile(id string) error
	SwitchPublic(id string) error
}

type FileRepositoryImpl struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) *FileRepositoryImpl {
	return &FileRepositoryImpl{db: db}
}

func (u *FileRepositoryImpl) CreateFile(file models.File) error {
	query := "INSERT INTO files (file_name, file_path, user_id, is_public, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	_, err := u.db.Exec(query, file.FileName, file.FilePath, file.UserId, file.IsPublic, createdAt, updatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (u *FileRepositoryImpl) GetFile(id string) (models.File, error) {
	query := "SELECT * FROM files WHERE id = $1 AND deleted_at IS NULL"
	err := u.db.Get(&file, query, id)
	if err != nil {
		return file, err
	}
	return file, nil
}

func (u *FileRepositoryImpl) GetFiles() ([]models.File, error) {
	query := "SELECT * FROM files WHERE deleted_at IS NULL"
	var files []models.File
	err := u.db.Select(&files, query)
	if err != nil {
		return files, err
	}
	return files, nil
}

func (u *FileRepositoryImpl) UpdateFile(file models.File, id string) error {
	query := "UPDATE files SET file_name = $1, is_public = $2, updated_at = $2 WHERE id = $3"
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	_, err := u.db.Exec(query, file.FileName, file.IsPublic, updatedAt, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *FileRepositoryImpl) DeleteFile(id string) error {
	query := "UPDATE files SET deleted_at = $1 WHERE id = $2"
	deletedAt := time.Now().Format("2006-01-02 15:04:05")
	_, err := u.db.Exec(query, deletedAt, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *FileRepositoryImpl) SwitchPublic(id string) error {
	query := "SELECT * FROM files WHERE id = $1 AND deleted_at IS NULL"
	err := u.db.Get(&file, query, id)
	if err != nil {
		return err
	}

	query2 := "UPDATE files SET is_public = $1 WHERE id = $2"

	_, err = u.db.Exec(query2, !file.IsPublic, id)
	if err != nil {
		return err
	}
	return nil
}
