package main

import (
	"fmt"
	"net/http"

	"github.com/phainosz/golang-crud/internal/config"
	"github.com/phainosz/golang-crud/internal/router"
)

func main() {
	config.LoadEnvironmentVariables()

	router := router.NewRouter()

	fmt.Printf("Server started on port %d\n", config.ServerPort)
	http.ListenAndServe(fmt.Sprintf(":%d", config.ServerPort), router)
}
