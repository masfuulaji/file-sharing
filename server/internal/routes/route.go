package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/masfuulaji/file-sharing/internal/app/handler"
	"github.com/masfuulaji/file-sharing/internal/database"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func SetupRoutes(r *chi.Mux) {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
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

	userHandler := handler.NewUserHandler(db.DB)
	r.Route("/user", func(r chi.Router) {
		r.Get("/", userHandler.GetUsers)
		r.Get("/{id}", userHandler.GetUser)
		r.Post("/", userHandler.CreateUser)
		r.Put("/{id}", userHandler.UpdateUser)
		r.Delete("/{id}", userHandler.DeleteUser)
	})
}
