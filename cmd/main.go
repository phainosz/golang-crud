package main

import (
	"log"

	"github.com/phainosz/golang-crud/pkg/db"
)

func main() {
	db, err := db.Connect()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	statement, err := db.Prepare("insert into users (name, email) values (?, ?)")
	if err != nil {
		log.Fatal("Error preparing statement")
	}
	defer statement.Close()

	_, err = statement.Exec("test name", "test@gmail.com")
	if err != nil {
		log.Fatal("Error executing statement")
	}
}
