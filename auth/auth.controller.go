package auth

import (
	"portofolio-api/auth/services"

	"github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine) {
	authRouter := r.Group("/auth")

	authRouter.POST("/register", services.CreateUser)
	authRouter.POST("/login", services.LoginUser)
}
