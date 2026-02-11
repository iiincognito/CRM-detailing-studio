package db

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConn := os.Getenv("CONN_DB")
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Println("Error pinging database")
		db.Close()
		return nil, err
	}
	return db, nil
}
