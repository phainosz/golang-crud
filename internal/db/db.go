package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/phainosz/golang-crud/internal/config"
)

// Creates the database connection using mysql
func Connect() (*sql.DB, error) {
	connectionUrl := config.GetDbConnection()
	db, err := sql.Open("mysql", connectionUrl)

	//here connection is ok
	if err != nil {
		return nil, err
	}

	//here checks if user is correct after connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("database connected")

	return db, nil
}
