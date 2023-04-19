package photo

import "github.com/gin-gonic/gin"

type PhotoHandler interface {
	FindAllPhotosHdl(ctx *gin.Context)
	FindPhotoByIdHdl(ctx *gin.Context)
	CreatePhotoHdl(ctx *gin.Context)
	UpdatePhotoHdl(ctx *gin.Context)
	DeletePhotoByIdHdl(ctx *gin.Context)
}