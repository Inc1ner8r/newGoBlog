package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConnectDb() *sql.DB {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	DbString := username + ":" + password + "@tcp(127.0.0.1:3306)/" + dbName + "?parseTime=true"

	if db, err := sql.Open("mysql", DbString); err != nil {
		panic(err.Error())
	} else {
		defer db.Close()
		return db
	}
}
