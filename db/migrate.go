package db

import (
	// "database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	// _ "github.com/lib/pq"
	// "github.com/golang-migrate/migrate/v4/database/postgres"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	// "os"
)

// Up apply all DB migrations
func Up(driver database.Driver) {
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err != nil {
		log.Println(err)
	}
	if err := m.Up(); err != nil {
		log.Println(err)
	}
	log.Println("DB UP")
}
