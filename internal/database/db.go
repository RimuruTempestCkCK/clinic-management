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

	// Migrate schemas
	err = db.AutoMigrate(
		&User{},
		&Patient{},
		&Doctor{},
		&Specialization{},
		&Appointment{},
		&MedicalRecord{},
		&Medicine{},
		&Prescription{},
		&PrescriptionItem{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database. \n", err)
	}

	DB = db
}
