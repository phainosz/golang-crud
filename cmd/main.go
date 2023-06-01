package main

import (
	"fmt"
	"log"

	"github.com/phainosz/golang-crud/internal/config"
	"github.com/phainosz/golang-crud/internal/models"
	"github.com/phainosz/golang-crud/internal/repositories"
)

func main() {
	config.LoadEnvironmentVariables()

	repositories.CreateUser(models.User{Name: "User", Email: "user@gmail.com"})

	users, err := repositories.GetUsers()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)

	if err := repositories.DeleteUserById(6); err != nil {
		log.Fatal(err)
	}
}
