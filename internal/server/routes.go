package server

import (
	handler "dockeeter-API/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initializeRoutes() *http.Server {
	r := chi.NewRouter()

	r.Get("/containers", handler.GetContainers)

	return &http.Server{
		Addr:    "3000",
		Handler: r,
	}
}
