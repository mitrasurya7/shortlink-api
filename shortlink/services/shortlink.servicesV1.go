package services

import (
	"log"
	"net/http"
	"os"
	"shortlink-api/database"
	"shortlink-api/helpers"
	"shortlink-api/shortlink/entities"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CreateShortlink(c *gin.Context) {
	db := database.OpenConnection()

	shortlinkInput := &entities.Shortlink_tab{}
	if err := c.ShouldBindJSON(shortlinkInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idToken := helpers.GenerateRandomString(4)

	for {
		var existingShortlink entities.Shortlink_tab
		err := db.Where("Shortlink = ?", idToken).First(&existingShortlink).Error
		if err != nil {
			break
		}
		idToken = helpers.GenerateRandomString(4)
	}

	shortlink := &entities.Shortlink_tab{
		Shortlink:    idToken,
		Redirectlink: shortlinkInput.Redirectlink,
	}

	if err := db.Create(&shortlink).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	frondendUrl := os.Getenv("FRONTEND_URL")

	c.JSON(http.StatusCreated, gin.H{
		"shortlink": frondendUrl + shortlink.Shortlink,
		"message":   "Shortlink berhasil dibuat",
	})
}

func GetShortlinkByShortId(c *gin.Context) {
	db := database.OpenConnection()

	shortlinkID := c.Param("shortlinkId")

	shortlink := &entities.Shortlink_tab{}
	if err := db.Where("Shortlink = ?", shortlinkID).First(&shortlink).Error; err != nil {
		c.JSON(404, gin.H{"error": "Shortlink tidak ditemukan"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, shortlink.Redirectlink)
}
