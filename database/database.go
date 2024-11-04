package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDatbase() {
	dsn := GlobalConfig.DSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "project_manage.",
		},
	})
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	DB = db
}
