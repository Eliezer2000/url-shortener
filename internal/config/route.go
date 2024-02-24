package config

import (
	"crud-app/internal/controller"
	"crud-app/internal/server"
	"crud-app/internal/service"
	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux {
	urlService := service.NewUrlService()
	urlController := controller.NewUrlController(urlService)
	router := chi.NewRouter()
	router.Get("/Hello", server.TestHandler())
	router.Post("/url", urlController.StoreShortUrlHandler)
	router.Get("/url/{shortUrl}", urlController.GetShortURLHandler)
	return router
}