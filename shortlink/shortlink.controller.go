package shortlink

import (
	"portofolio-api/shortlink/services"

	"github.com/gin-gonic/gin"
)

func ControllerShortLink(r *gin.Engine) {
	rShortlink := r.Group("")
	rShortlink.POST("", services.CreateShortlink)
	rShortlink.GET(":shortlinkId", services.GetShortlinkByShortId)
}
