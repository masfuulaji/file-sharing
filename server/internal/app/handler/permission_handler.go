package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/file-sharing/internal/app/models"
	"github.com/masfuulaji/file-sharing/internal/app/repositories"
)

type PermissionHandler interface {
	CreatePermission(w http.ResponseWriter, r *http.Request)
	UpdatePermission(w http.ResponseWriter, r *http.Request)
	DeletePermission(w http.ResponseWriter, r *http.Request)
	GetPermission(w http.ResponseWriter, r *http.Request)
	GetPermissions(w http.ResponseWriter, r *http.Request)
	GetPermissionsByUser(w http.ResponseWriter, r *http.Request)
}

type PermissionHandlerImpl struct {
	permissionRepository *repositories.PermissionRepositoryImpl
}

func NewPermissionHandler(db *sqlx.DB) *PermissionHandlerImpl {
	return &PermissionHandlerImpl{permissionRepository: repositories.NewPermissionRepository(db)}
}

func (u *PermissionHandlerImpl) CreatePermission(w http.ResponseWriter, r *http.Request) {
	permission := models.Permission{}
	err := json.NewDecoder(r.Body).Decode(&permission)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	err = u.permissionRepository.CreatePermission(permission)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	response := Response{
		Message: "Permission created successfully",
		Status:  200,
	}
	json.NewEncoder(w).Encode(response)
}

func (u *PermissionHandlerImpl) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	permission := models.Permission{}
	err := json.NewDecoder(r.Body).Decode(&permission)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	err = u.permissionRepository.UpdatePermission(permission, id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	response := Response{
		Message: "Permission updated successfully",
		Status:  200,
	}
	json.NewEncoder(w).Encode(response)
}

func (u *PermissionHandlerImpl) DeletePermission(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := u.permissionRepository.DeletePermission(id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	response := Response{
		Message: "Permission deleted successfully",
		Status:  200,
	}
	json.NewEncoder(w).Encode(response)
}

func (u *PermissionHandlerImpl) GetPermission(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	permission, err := u.permissionRepository.GetPermission(id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(permission)
}

func (u *PermissionHandlerImpl) GetPermissions(w http.ResponseWriter, r *http.Request) {
	permissions, err := u.permissionRepository.GetPermissions()
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(permissions)
}

func (u *PermissionHandlerImpl) GetPermissionsByUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	permissions, err := u.permissionRepository.GetPermissionsByUser(id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(permissions)
}
