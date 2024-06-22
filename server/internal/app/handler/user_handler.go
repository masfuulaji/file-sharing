package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/file-sharing/internal/app/models"
	"github.com/masfuulaji/file-sharing/internal/app/repositories"
)

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
}

type UserHandlerImpl struct {
	userRepository *repositories.UserRepositoryImpl
}

func NewUserHandler(db *sqlx.DB) *UserHandlerImpl {
	return &UserHandlerImpl{userRepository: repositories.NewUserRepository(db)}
}

func (u *UserHandlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	err = u.userRepository.CreateUser(user)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	response := Response{
		Message: "User created successfully",
		Status:  200,
	}
	json.NewEncoder(w).Encode(response)
}

func (u *UserHandlerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	err = u.userRepository.UpdateUser(user, id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	response := Response{
		Message: "User updated successfully",
		Status:  200,
	}
	json.NewEncoder(w).Encode(response)
}

func (u *UserHandlerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := u.userRepository.DeleteUser(id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	response := Response{
		Message: "User deleted successfully",
		Status:  200,
	}
	json.NewEncoder(w).Encode(response)
}

func (u *UserHandlerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := u.userRepository.GetUser(id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (u *UserHandlerImpl) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.userRepository.GetUsers()
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(users)
}
