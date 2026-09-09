package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Default fallback for local testing
		dsn = "host=localhost user=postgres password=postgres dbname=clinic port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("Database connected successfully")

	// Schema creation is handled manually via seed_with_schema.sql
	// db.AutoMigrate(...) is removed to prevent constraint conflicts.

	DB = db
}
