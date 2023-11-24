package main

import (
	"fmt"
	"log"
	"os"
	"shortlink-api/database"
	"shortlink-api/shortlink"

	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("APP_PORT")
	frondendUrl := os.Getenv("ORIGIN_URL")

	fmt.Println("Server running on port", port)
	database.AutoMigration()
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{frondendUrl, "http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))
	shortlink.ControllerShortLink(r)

	r.Run(":" + port)
}
