package main

import (
	"fmt"
	"log"
	"os"
	"shortlink-api/database"
	"shortlink-api/shortlink"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("APP_PORT")

	fmt.Println("Server running on port", port)
	database.AutoMigration()
	r := gin.Default()
	shortlink.ControllerShortLink(r)

	r.Run(":" + port)
}
