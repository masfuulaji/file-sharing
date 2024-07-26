package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/file-sharing/internal/app/models"
)

type PermissionRepository interface {
	CreatePermission(permission models.Permission) error
	GetPermission(id string) (models.Permission, error)
	GetPermissions() ([]models.Permission, error)
	UpdatePermission(permission models.Permission, id string) error
	DeletePermission(id string) error
	GetPermissionsByUser(userId string) ([]models.Permission, error)
}

type PermissionRepositoryImpl struct {
	db *sqlx.DB
}

func NewPermissionRepository(db *sqlx.DB) *PermissionRepositoryImpl {
	return &PermissionRepositoryImpl{db: db}
}

func (u *PermissionRepositoryImpl) CreatePermission(permission models.Permission) error {
	query := "INSERT INTO permissions (user_id, file_id, permission, created_at) VALUES ($1, $2, $3, $4)"
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	_, err := u.db.Exec(query, permission.UserId, permission.FileId, permission.Permission, createdAt)
	if err != nil {
		return err
	}
	return nil
}

func (u *PermissionRepositoryImpl) GetPermission(id string) (models.Permission, error) {
	var permission models.Permission
	query := "SELECT * FROM permissions WHERE id = $1"
	err := u.db.Get(&permission, query, id)
	if err != nil {
		return permission, err
	}
	return permission, nil
}

func (u *PermissionRepositoryImpl) GetPermissions() ([]models.Permission, error) {
	var permissions []models.Permission
	query := "SELECT * FROM permissions"
	err := u.db.Select(&permissions, query)
	if err != nil {
		return permissions, err
	}
	return permissions, nil
}

func (u *PermissionRepositoryImpl) UpdatePermission(permission models.Permission, id string) error {
	query := "UPDATE permissions SET user_id = $1, file_id = $2 WHERE id = $3"
	_, err := u.db.Exec(query, permission.UserId, permission.FileId, permission.Permission, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *PermissionRepositoryImpl) DeletePermission(id string) error {
	query := "DELETE FROM permissions WHERE id = $1"
	_, err := u.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *PermissionRepositoryImpl) GetPermissionsByUser(userId string) ([]models.Permission, error) {
	var permissions []models.Permission
	query := "SELECT * FROM permissions WHERE user_id = $1"
	err := u.db.Select(&permissions, query, userId)
	if err != nil {
		return permissions, err
	}
	return permissions, nil
}
