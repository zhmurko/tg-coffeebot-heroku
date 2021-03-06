package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	// "github.com/golang-migrate/migrate/v4"
	// "github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
)

var DB *sql.DB

func Connect() *sql.DB {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		url = "postgres://postgres@localhost:5432/coffee?sslmode=disable"
	}

	db, err := sql.Open("postgres", url)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	Up(driver)
	if err != nil {
		log.Println(err)
	}
	return db
}
