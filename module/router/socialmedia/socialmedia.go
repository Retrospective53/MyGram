package socialmedia

import (
	socialmediahandler "github.com/Retrospective53/myGram/module/handler/socialmedia"
	"github.com/Retrospective53/myGram/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func NewSocialMediaRouter(v1 *gin.RouterGroup, socialMediaHdl socialmediahandler.SocialMediaHandler) {
	g := v1.Group("/socialmedias")

	// register all router
	g.GET("/:id", middleware.BearerOAuth(), socialMediaHdl.FindSocialMediaByIdHdl)
	g.GET("/all", middleware.BearerOAuth(), socialMediaHdl.FindAllSocialMediasHdl)
	g.POST("", middleware.BearerOAuth(), socialMediaHdl.CreateSocialMediaHdl)
	g.PUT("/:id", middleware.BearerOAuth(), socialMediaHdl.UpdateSocialMediaHdl)
	g.DELETE("/:id", middleware.BearerOAuth(), socialMediaHdl.DeleteSocialMediaByIdHdl)
}