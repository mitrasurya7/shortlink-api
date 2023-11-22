package services

import (
	"net/http"
	"portofolio-api/auth/entities"
	"portofolio-api/database"
	"portofolio-api/middlewares"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	db := database.OpenConnection()
	userInput := &entities.User_tab{}

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashPassword, err := middlewares.EncryptPassword(userInput.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi kata sandi"})
		return
	}

	db.Create(&entities.User_tab{
		Username: userInput.Username,
		Password: hashPassword,
	})

	c.JSON(http.StatusAccepted, gin.H{

		"message": "user success created",
	})

}

type FormLogin struct {
	Username string
	Password string
}

func LoginUser(c *gin.Context) {
	db := database.OpenConnection()
	var formLogin FormLogin

	err := c.ShouldBindJSON(&formLogin)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	userData := &entities.User_tab{}

	if err := db.Where("Username = ?", formLogin.Username).First(&userData).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong email or password",
		})
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(formLogin.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong email or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Success",
	})
}
