package comment

import (
	commenthandler "github.com/Retrospective53/myGram/module/handler/comment"
	"github.com/Retrospective53/myGram/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func NewCommentRouter(v1 *gin.RouterGroup, commentHdl commenthandler.CommentHandler) {
	g := v1.Group("/comments")

	// register all router
	g.GET("/:id", middleware.BearerOAuth(), commentHdl.FindCommentByIdHdl)
	g.GET("/all", middleware.BearerOAuth(), commentHdl.FindAllCommentsHdl)
	g.POST("", middleware.BearerOAuth(), commentHdl.CreateCommentHdl)
	g.PUT("/:id", middleware.BearerOAuth(), commentHdl.UpdateCommentHdl)
	g.DELETE("/:id", middleware.BearerOAuth(), commentHdl.DeleteCommentByIdHdl)
}