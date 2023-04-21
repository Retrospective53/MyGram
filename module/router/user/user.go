package user

import (
	accounthandler "github.com/Retrospective53/myGram/module/handler/user"
	"github.com/Retrospective53/myGram/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func NewAccountRouter(v1 *gin.RouterGroup, accountHdl accounthandler.UserHandler) {
	g := v1.Group("/accounts")

	// register all router
	g.POST("",
		accountHdl.CreateAccount)
	g.POST("/login",
		accountHdl.LoginAccount)
	g.GET("",
		middleware.BearerOAuth(),
		accountHdl.GetAccount)
}