package database

import (
	"log"
	"os"
	"shortlink-api/shortlink/entities"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dialect := postgres.Open("host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Jakarta")
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic("failed Connecting database")
	}
	return db
}

func AutoMigration() {
	db := OpenConnection()
	db.AutoMigrate(&entities.Shortlink_tab{})
}
