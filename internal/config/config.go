package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DbConnection = ""
	ServerPort   = 0
)

// loads all environment variables
func LoadEnvironmentVariables() {
	var err error
	if err = godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	DbConnection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	ServerPort, err = strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		ServerPort = 8000
	}
}
