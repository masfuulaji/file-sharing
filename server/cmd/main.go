package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/masfuulaji/file-sharing/internal/routes"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	routes.SetupRoutes(r)
	http.ListenAndServe(":3000", r)
}
