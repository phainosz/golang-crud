package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Route struct {
	Uri      string
	Method   string
	Function func(w http.ResponseWriter, r *http.Request)
}

// add all routes to router
func Config(r *chi.Mux) {
	routes := usersRoutes
	//if need more routes, follow the same as users and append
	//routes = append(routes, otherRoutes)

	for _, route := range routes {
		r.MethodFunc(route.Method, route.Uri, route.Function)
	}
}
