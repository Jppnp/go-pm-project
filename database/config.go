package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	Database struct {
		Driver   string `json:"driver"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
		SSLMode  string `json:"sslmode"`
		Timezone string `json:"timezone"`
	} `json:"database"`
}

func (c *Config) DSN() string {
	var dsn strings.Builder
	// dsn.WriteString(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
	// 	c.Database.Host,
	// 	c.Database.Port,
	// 	c.Database.User,
	// 	c.Database.Password,
	// 	c.Database.DBName,
	// 	c.Database.SSLMode,
	// 	c.Database.Timezone,
	// ))
	dsn.WriteString(fmt.Sprintf("user=%s dbname=%s sslmode=%s", c.Database.User, c.Database.DBName, c.Database.SSLMode))
	return dsn.String()
}

func loadConfig() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func ConnectDatabase() {
	config, err := loadConfig()
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	dsn := config.DSN()
	fmt.Printf("Using DSN: %s\n", dsn)

	// Connect to the database
	db, err := sql.Open(config.Database.Driver, dsn)
	if err != nil {
		log.Fatal("failed to connect to the database for migration:", err)
	}
	defer db.Close()

	// Check the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal("failed to ping the database:", err)
	}
	fmt.Println("Database connection established successfully.")

	// Print current migration version
	var currentVersion int
	err = db.QueryRow("SELECT version_id FROM public.goose_db_version ORDER BY version_id DESC LIMIT 1").Scan(&currentVersion)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal("failed to query current migration version:", err)
	}
	fmt.Printf("Current migration version: %d\n", currentVersion)

	// Check where the `goose_db_version` table is located
	rows, err := db.Query(`
		SELECT table_schema, table_name
		FROM information_schema.tables
		WHERE table_name = 'goose_db_version';
	`)
	if err != nil {
		log.Fatal("failed to query information_schema.tables:", err)
	}
	defer rows.Close()

	fmt.Println("Table `goose_db_version` is located in the following schema(s):")
	for rows.Next() {
		var schema, table string
		if err := rows.Scan(&schema, &table); err != nil {
			log.Fatal("failed to scan row:", err)
		}
		fmt.Printf("Schema: %s, Table: %s\n", schema, table)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("error occurred during rows iteration:", err)
	}

	// Run migrations
	goose.SetBaseFS(os.DirFS("migrations"))
	fmt.Println("Starting Goose migration...")
	if err := goose.Up(db, "."); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("Goose migration completed successfully.")

	// Initialize GORM
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	DB = gormDB
}
