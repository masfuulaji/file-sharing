package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
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

	dst, err := os.Create(filepath.Join("storage", header.Filename))
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
	w.Write([]byte("Success"))
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	fileId := chi.URLParam(r, "file")
	if fileId == "" {
		w.Write([]byte("File id is required"))
		return
	}
	file, err := os.ReadFile("storage/" + fileId)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(file)
}
