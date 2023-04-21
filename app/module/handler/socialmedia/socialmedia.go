package socialmedia

import "github.com/gin-gonic/gin"

type SocialMediaHandler interface {
	FindAllSocialMediasHdl(ctx *gin.Context)
	FindSocialMediaByIdHdl(ctx *gin.Context)
	CreateSocialMediaHdl(ctx *gin.Context)
	UpdateSocialMediaHdl(ctx *gin.Context)
	DeleteSocialMediaByIdHdl(ctx *gin.Context)
}