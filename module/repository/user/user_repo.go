package user

import (
	"context"

	"github.com/Retrospective53/myGram/module/models"
)

type UserRepo interface {
	CreateAccount(ctx context.Context, acc models.User) (created models.User, err error) 
	GetAccountByUserName(ctx context.Context, username string) (account models.User, err error) 
	GetAccountByUserId(ctx context.Context, userId string) (account models.User, err error) 
}