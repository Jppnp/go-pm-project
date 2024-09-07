package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/pressly/goose/v3"
)

func Migration() {
	// connStr := "user=postgres dbname=postgres-golang sslmode=disable"
	config, _ := loadConfig()
	dsn := config.DSN()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open a DB connection: %v", err)
	}
	defer db.Close()

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Database migration completed successfully")
}
