package photo

import (
	photohandler "github.com/Retrospective53/myGram/module/handler/photo"
	"github.com/Retrospective53/myGram/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func NewPhotoRouter(v1 *gin.RouterGroup, photoHdl photohandler.PhotoHandler) {
	g := v1.Group("/photo")

	// register all router
	g.GET("/:id", middleware.BearerOAuth(), photoHdl.FindPhotoByIdHdl)
	g.GET("/all", middleware.BearerOAuth(), photoHdl.FindAllPhotosHdl)
	g.POST("", middleware.BearerOAuth(), photoHdl.CreatePhotoHdl)
	g.PUT("/:id", middleware.BearerOAuth(), photoHdl.UpdatePhotoHdl)
	g.DELETE("/:id", middleware.BearerOAuth(), photoHdl.DeletePhotoByIdHdl)
}