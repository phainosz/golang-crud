package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Create the connection with database using mysql
func Connect() *sql.DB {
	connectionUrl := "myuser:mypassword@tcp(localhost:3306)/golang-crud?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", connectionUrl)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("database connected")

	return db
}
