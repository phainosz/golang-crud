package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/phainosz/golang-crud/internal/config"
)

// Creates the database connection using mysql
func Connect() (*sql.DB, error) {
	connectionUrl := config.DbConnection
	db, err := sql.Open("mysql", connectionUrl)

	//check if connection is ok
	if err != nil {
		return nil, err
	}

	//checks if user and password is correct after connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("database connected")

	return db, nil
}
