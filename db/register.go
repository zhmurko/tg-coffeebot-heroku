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

// connection instance
var DB *sql.DB

// inititate a connection to PostgreSQL service
func Connect() *sql.DB {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		url = "postgres://postgres@localhost:5432/coffee?sslmode=disable"
	}

	db, _ := sql.Open("postgres", url)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	Up(driver)
	if err != nil {
		log.Println(err)
	}
	return db
}
