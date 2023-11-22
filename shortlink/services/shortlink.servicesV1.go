package services

import (
	"fmt"
	"net/http"
	"portofolio-api/database"
	"portofolio-api/shortlink/entities"

	"github.com/gin-gonic/gin"
)

func CreateShortlink(c *gin.Context) {
	db := database.OpenConnection()
	shortlinkInput := &entities.Shortlink_tab{}

	err := c.ShouldBindJSON(&shortlinkInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(shortlinkInput)
	fmt.Println(shortlinkInput)

	c.JSON(http.StatusCreated, gin.H{
		"shortlinkId": shortlinkInput.Shortlink,
		"redirectUrl": shortlinkInput.Redirectlink,
	})

}

func GetShortlinkByShortId(c *gin.Context) {
	db := database.OpenConnection()

	shortlinkID := c.Param("shortlinkId")

	shortlink := &entities.Shortlink_tab{}
	if err := db.Where("Shortlink = ?", shortlinkID).First(&shortlink).Error; err != nil {
		c.JSON(404, gin.H{"error": "Shortlink tidak ditemukan"}) // Atau respon sesuai kebutuhan Anda
		return
	}

	c.Redirect(http.StatusMovedPermanently, shortlink.Redirectlink)
}
