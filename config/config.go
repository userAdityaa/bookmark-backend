package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/userAdityaa/bookmark-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	err := godotenv.Load()
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := "host=localhost user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database")
		return nil, err
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	log.Println("Database connected")
	return db, nil
}

func DisconnectDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Failed to close the database connection:", err)
		return
	}
	sqlDB.Close()
}
