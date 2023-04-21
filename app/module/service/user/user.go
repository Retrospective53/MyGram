package user

import (
	"context"

	accountmodel "github.com/Retrospective53/myGram/module/models"
	token "github.com/Retrospective53/myGram/module/models/token"
)

type UserService interface {
	CreateAccountSvc(ctx context.Context, acc accountmodel.CreateAccount) (created accountmodel.AccountResponse, err error)
	LoginAccountByUserNameSvc(ctx context.Context, loginAcc accountmodel.LoginAccount) (tokens token.Tokens, err error)
	GetAccountSvc(ctx context.Context, userId string) (account accountmodel.AccountResponse, err error)
}