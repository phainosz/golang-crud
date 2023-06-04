package main

import (
	"fmt"
	"log"

	"github.com/phainosz/golang-crud/internal/config"
	"github.com/phainosz/golang-crud/internal/repositories"
)

func main() {
	config.LoadEnvironmentVariables()

	users, err := repositories.GetUsers()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)
}
