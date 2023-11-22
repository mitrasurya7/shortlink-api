package main

import (
	"portofolio-api/auth"
	"portofolio-api/database"
	"portofolio-api/shortlink"

	"github.com/gin-gonic/gin"
)

func main() {
	database.AutoMigration()
	r := gin.Default()
	shortlink.ControllerShortLink(r)
	auth.Auth(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
