package services

import (
	"net/http"
	"shortlink-api/database"
	"shortlink-api/helpers"
	"shortlink-api/shortlink/entities"

	"github.com/gin-gonic/gin"
)

func CreateShortlink(c *gin.Context) {
	db := database.OpenConnection()
	shortlinkInput := &entities.Shortlink_tab{}
	if err := c.ShouldBindJSON(shortlinkInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idToken := helpers.GenerateRandomString(4)

	shortlink := &entities.Shortlink_tab{
		Shortlink:    idToken,
		Redirectlink: shortlinkInput.Redirectlink,
	}

	readyToken := db.Where("Shortlink = ?", idToken).First(&shortlink).Error
	for readyToken == nil {
		idToken = helpers.GenerateRandomString(4)
	}

	if err := db.Create(&shortlink).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shortlink": shortlink,
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
