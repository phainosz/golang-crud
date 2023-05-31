package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Creates the database connection using mysql
func Connect() (*sql.DB, error) {
	connectionUrl := "user:123@tcp(localhost:3306)/golang_crud?charset=utf8&parseTime=True&loc=Local"
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
