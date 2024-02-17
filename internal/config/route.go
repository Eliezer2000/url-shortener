package config

import (
	"crud-app/internal/server"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Get("Hello", server.TestHandler())
	return router
}