package shortlink

import (
	"shortlink-api/shortlink/services"

	"github.com/gin-gonic/gin"
)

func ControllerShortLink(r *gin.Engine) {
	rShortlink := r.Group("")
	rShortlink.POST("/short", services.CreateShortlink)
	rShortlink.GET(":shortlinkId", services.GetShortlinkByShortId)
}
