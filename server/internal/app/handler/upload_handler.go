package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/file-sharing/internal/app/models"
	"github.com/masfuulaji/file-sharing/internal/app/repositories"
	"github.com/masfuulaji/file-sharing/internal/utils"
)

type UploadHandler interface {
	UploadFile(w http.ResponseWriter, r *http.Request)
	DownloadHandler(w http.ResponseWriter, r *http.Request)
}

type UploadHandlerImpl struct {
	fileRepository *repositories.FileRepositoryImpl
}

func NewUploadHandler(db *sqlx.DB) *UploadHandlerImpl {
	return &UploadHandlerImpl{fileRepository: repositories.NewFileRepository(db)}
}

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (u *UploadHandlerImpl) UploadFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		response := Response{
			Message: err.Error(),
			Status:  500,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		response := Response{
			Message: err.Error(),
			Status:  500,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	defer file.Close()

	randString := utils.RandomString(10)
	newFileName := fmt.Sprintf("%s%s", randString, filepath.Ext(header.Filename))
	dst, err := os.Create(filepath.Join("storage", newFileName))
	if err != nil {
		response := Response{
			Message: err.Error(),
			Status:  500,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		response := Response{
			Message: err.Error(),
			Status:  500,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	cookies, err := r.Cookie("token")
	if err != nil {
		response := Response{
			Message: err.Error(),
			Status:  500,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	tokenStr := cookies.Value
	claims := &Claims{}
	_, err = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil {
		response := Response{
			Message: err.Error(),
			Status:  500,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	userId := claims.Id

	err = u.fileRepository.CreateFile(models.File{FileName: header.Filename, FilePath: newFileName, IsPublic: false, UserId: strconv.Itoa(userId)})
	if err != nil {
		response := Response{
			Message: err.Error(),
			Status:  500,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Write([]byte("Success"))
}

func (u *UploadHandlerImpl) DownloadHandler(w http.ResponseWriter, r *http.Request) {
	cookies, err := r.Cookie("token")
	if err != nil {
		response := Response{
			Message: err.Error(),
			Status:  500,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	tokenStr := cookies.Value
	claims := &Claims{}
	_, err = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil {
		response := Response{
			Message: err.Error(),
			Status:  500,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	fileId := chi.URLParam(r, "id")

	file, err := u.fileRepository.GetFile(fileId)
	if err != nil {
		response := Response{
			Message: err.Error(),
			Status:  500,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if file.UserId != strconv.Itoa(claims.Id) && !file.IsPublic {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	fileDownload, err := os.ReadFile("storage/" + file.FilePath)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+file.FileName)
	w.Write(fileDownload)
}
