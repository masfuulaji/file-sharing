package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/masfuulaji/file-sharing/internal/app/handler"
	"github.com/masfuulaji/file-sharing/internal/database"
	"github.com/masfuulaji/file-sharing/internal/utils"
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

	loginHandler := handler.NewLoginHandler(db.DB)
	r.Post("/login", loginHandler.Login)
	r.With(utils.AuthMiddleware).Get("/protected", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Protected!"))
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
