package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/file-sharing/internal/app/models"
	"github.com/masfuulaji/file-sharing/internal/app/repositories"
)

type LoginHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type LoginHandlerImpl struct {
	userRepository *repositories.UserRepositoryImpl
}

func NewLoginHandler(db *sqlx.DB) *LoginHandlerImpl {
	return &LoginHandlerImpl{userRepository: repositories.NewUserRepository(db)}
}

func (u *LoginHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	res, err := u.userRepository.GetUserByUsername(user.Username)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if user.Password != res.Password {
		json.NewEncoder(w).Encode("Incorrect password")
		return
	}

	json.NewEncoder(w).Encode(user)
}
