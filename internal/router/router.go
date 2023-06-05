package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/phainosz/golang-crud/internal/router/routes"
)

// create and config router
func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Heartbeat("/"))

	routes.Config(router)
	return router
}
