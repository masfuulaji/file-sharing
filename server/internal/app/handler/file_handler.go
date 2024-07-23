package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/file-sharing/internal/app/models"
	"github.com/masfuulaji/file-sharing/internal/app/repositories"
)

type FileHandler interface {
	CreateFile(w http.ResponseWriter, r *http.Request)
	UpdateFile(w http.ResponseWriter, r *http.Request)
	DeleteFile(w http.ResponseWriter, r *http.Request)
	GetFile(w http.ResponseWriter, r *http.Request)
	GetFiles(w http.ResponseWriter, r *http.Request)
}

type FileHandlerImpl struct {
	fileRepository *repositories.FileRepositoryImpl
}

func NewFileHandler(db *sqlx.DB) *FileHandlerImpl {
	return &FileHandlerImpl{fileRepository: repositories.NewFileRepository(db)}
}

func (f *FileHandlerImpl) CreateFile(w http.ResponseWriter, r *http.Request) {
	file := models.File{}
	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	err = f.fileRepository.CreateFile(file)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	response := Response{
		Message: "File created successfully",
		Status:  200,
	}
	json.NewEncoder(w).Encode(response)
}

func (f *FileHandlerImpl) UpdateFile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	file := models.File{}
	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	err = f.fileRepository.UpdateFile(file, id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	response := Response{
		Message: "File updated successfully",
		Status:  200,
	}
	json.NewEncoder(w).Encode(response)
}

func (f *FileHandlerImpl) DeleteFile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := f.fileRepository.DeleteFile(id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	response := Response{
		Message: "File deleted successfully",
		Status:  200,
	}
	json.NewEncoder(w).Encode(response)
}

func (f *FileHandlerImpl) GetFile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	file, err := f.fileRepository.GetFile(id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(file)
}

func (f *FileHandlerImpl) GetFiles(w http.ResponseWriter, r *http.Request) {
	files, err := f.fileRepository.GetFiles()
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(files)
}
