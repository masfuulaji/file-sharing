package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/masfuulaji/file-sharing/internal/app/handler"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func SetupRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Ell!"))
	})

	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Message: "Allo",
			Status:  200,
		}
		json.NewEncoder(w).Encode(response)
	})

	r.Post("/upload", handler.UploadHandler)
	r.Get("/download/{file}", handler.DownloadHandler)
}
