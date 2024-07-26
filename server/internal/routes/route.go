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
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Welcome"})
	})

	loginHandler := handler.NewLoginHandler(db.DB)
	r.Post("/login", loginHandler.Login)
	r.With(utils.AuthMiddleware).Get("/isLogin", loginHandler.IsLogin)

	uploadHandler := handler.NewUploadHandler(db.DB)
	r.Post("/upload", uploadHandler.UploadFile)
	r.Get("/download/{id}", uploadHandler.DownloadHandler)

	userHandler := handler.NewUserHandler(db.DB)
	r.Route("/user", func(r chi.Router) {
		r.Get("/", userHandler.GetUsers)
		r.Get("/{id}", userHandler.GetUser)
		r.Post("/", userHandler.CreateUser)
		r.Put("/{id}", userHandler.UpdateUser)
		r.Delete("/{id}", userHandler.DeleteUser)
	})

	fileHandler := handler.NewFileHandler(db.DB)
	r.Route("/file", func(r chi.Router) {
		r.Get("/", fileHandler.GetFiles)
		r.Get("/{id}", fileHandler.GetFile)
		r.Post("/", fileHandler.CreateFile)
		r.Put("/{id}", fileHandler.UpdateFile)
		r.Delete("/{id}", fileHandler.DeleteFile)
		r.Put("/switch_public/{id}", fileHandler.SwitchPublic)
	})

	permissionHandler := handler.NewPermissionHandler(db.DB)
	r.Route("/permission", func(r chi.Router) {
		r.Get("/", permissionHandler.GetPermissions)
		r.Get("/{id}", permissionHandler.GetPermission)
		r.Post("/", permissionHandler.CreatePermission)
		r.Put("/{id}", permissionHandler.UpdatePermission)
		r.Delete("/{id}", permissionHandler.DeletePermission)
		r.Get("/user/{id}", permissionHandler.GetPermissionsByUser)
	})
}
