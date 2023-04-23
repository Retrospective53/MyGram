package user

import "github.com/gin-gonic/gin"

type UserHandler interface {
	LoginAccount(ctx *gin.Context)
	CreateAccount(ctx *gin.Context)
	GetAccount(ctx *gin.Context)
}