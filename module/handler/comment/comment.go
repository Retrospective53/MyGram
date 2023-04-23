package comment

import "github.com/gin-gonic/gin"

type CommentHandler interface {
	FindAllCommentsHdl(ctx *gin.Context)
	FindCommentByIdHdl(ctx *gin.Context)
	CreateCommentHdl(ctx *gin.Context)
	UpdateCommentHdl(ctx *gin.Context)
	DeleteCommentByIdHdl(ctx *gin.Context)
}