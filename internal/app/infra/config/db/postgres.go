package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func CreatePostgresConnection() *sql.DB {
	godotenv.Load()
	connStr := os.Getenv("DB_URL")

	if len(connStr) == 0 {
		log.Fatal("connection string DB_URL is empty")
	}
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
