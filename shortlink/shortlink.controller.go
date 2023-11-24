package shortlink

import (
	"shortlink-api/shortlink/services"

	"github.com/gin-gonic/gin"
)

func ControllerShortLink(r *gin.Engine) {
	rShortlink := r.Group("")
	rShortlink.GET("/:shortlinkId", services.GetShortlinkByShortId)

	rShortlinkV1 := r.Group("/api/v1")
	rShortlinkV1.POST("/short", services.CreateShortlink)
}
