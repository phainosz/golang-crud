package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/phainosz/golang-crud/internal/router/routes"
)

// create and config router
func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Heartbeat("/"))
	router.Use(recoverer) //best to use middleware.Recoverer

	routes.Config(router)
	return router
}

// recover function to handle panic errors
func recoverer(handler http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()

			if err != nil {
				log.Println("Error: ", err)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)

				response := `{"error": "unexpected error."}`

				w.Write([]byte(response))
			}
		}()
		handler.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
