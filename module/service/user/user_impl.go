package user

import (
	"context"
	"sync"
	"time"

	"github.com/Retrospective53/myGram/module/models"
	token "github.com/Retrospective53/myGram/module/models/token"
	userrepo "github.com/Retrospective53/myGram/module/repository/user"
	crypto "github.com/Retrospective53/myGram/pkg/crypto"
	"github.com/google/uuid"
)

type UserServiceImpl struct {
	userRepo userrepo.UserRepo
}

func NewUserServiceImpl(userRepo userrepo.UserRepo) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}

func (a *UserServiceImpl) CreateAccountSvc(ctx context.Context, acc models.CreateAccount) (created models.AccountResponse, err error) {

	// need to hash password
	hashedPassowrd, err := crypto.GenerateHash(acc.Password)
	if err != nil {
		panic(err)
	}
	// update passowrd with hashed password
	acc.Password = hashedPassowrd
	// store to db
	createdAcc, err := a.userRepo.CreateAccount(ctx, models.User{
		ID:       uuid.New(),
		Username: acc.Username,
		Password: acc.Password,
		Role:     acc.Role,
		Email: acc.Email,
		Age: acc.Age,
	})
	// if err != nil {
	// 	panic(err)
	// }

	return models.AccountResponse{
		ID:        createdAcc.ID,
		Username:  createdAcc.Username,
		Role:      createdAcc.Role,
		CreatedAt: createdAcc.CreatedAt,
	}, err
}

func (a *UserServiceImpl) getAccountWithPassword(ctx context.Context, username string) (account models.AccountResponseWithPassword, err error) {
	// get account from database
	acc, err := a.userRepo.GetAccountByUserName(ctx, username)
	if err != nil {
		return
	}
	return models.AccountResponseWithPassword{
		AccountResponse: models.AccountResponse{
			ID:        acc.ID,
			Username:  acc.Username,
			Role:      acc.Role,
			CreatedAt: acc.CreatedAt,
		},
		Password: acc.Password,
	}, err
}

func (a *UserServiceImpl) generateAllTokensConcurrent(ctx context.Context, userid, username, role string) (idToken, accessToken, refreshToken string, err error) {
	// https://github.com/kataras/jwt
	timeNow := time.Now()
	defaultClaim := token.DefaultClaim{
		Expired:   int(timeNow.Add(24 * time.Hour).Unix()),
		NotBefore: int(timeNow.Unix()),
		IssuedAt:  int(timeNow.Unix()),
		Issuer:    "http://go-account",
		Audience:  "http://dts-07",
		Type:      token.ID_TOKEN,
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func(defaultClaim_ token.DefaultClaim) {
		defer wg.Done()
		// generate id token
		idTokenClaim := struct {
			token.DefaultClaim
			token.IDClaim
		}{
			DefaultClaim: defaultClaim_,
			IDClaim: token.IDClaim{
				Username: username,
				Role:     role,
			},
		}
		idToken, err = crypto.SignJWT(idTokenClaim)
		if err != nil {
			panic(err)
		}
	}(defaultClaim)

	go func(defaultClaim_ token.DefaultClaim) {
		defer wg.Done()
		// generate access token
		defaultClaim_.Expired = int(timeNow.Add(20 * time.Minute).UnixMilli())
		defaultClaim_.Type = token.ACCESS_TOKEN
		accessTokenClaim := struct {
			token.DefaultClaim
			token.AccessClaim
		}{
			DefaultClaim: defaultClaim_,
			AccessClaim: token.AccessClaim{
				Role:   role,
				UserID: userid,
			},
		}
		accessToken, err = crypto.SignJWT(accessTokenClaim)
		if err != nil {
			panic(err)
		}
	}(defaultClaim)

	go func(defaultClaim_ token.DefaultClaim) {
		defer wg.Done()
		// generate refresh token
		defaultClaim_.Expired = int(timeNow.Add(time.Hour).UnixMilli())
		defaultClaim_.Type = token.REFRESH_TOKEN
		refreshTokenClaim := struct {
			token.DefaultClaim
		}{
			DefaultClaim: defaultClaim_,
		}
		refreshToken, err = crypto.SignJWT(refreshTokenClaim)
		if err != nil {
			panic(err)
		}
	}(defaultClaim)

	wg.Wait()
	return
}



func (a *UserServiceImpl) LoginAccountByUserNameSvc(ctx context.Context, loginAcc models.LoginAccount) (tokens token.Tokens, err error) {
	// get account by username
	acc, err := a.getAccountWithPassword(ctx, loginAcc.Username)
	if err != nil {
		panic(err)
	}

	// compare password
	// password acc -> hashed password
	// password login acc -> plain password
	if err = crypto.CompareHash(acc.Password, loginAcc.Password); err != nil {
		panic(err)
	}

	// // record activity
	// createdActivity, err := a.activityRepo.CreateActivity(ctx, accountactivity.AccountActivity{
	// 	ID:     uuid.New(),
	// 	UserID: acc.ID,
	// 	Type:   accountactivity.ACTIVITY_LOGIN,
	// })
	// if err != nil {
	// 	logger.Error(ctx, "error when creating activity",
	// 		"logCtx", logCtx,
	// 		"error", err)
	// 	return
	// }

	idToken, accessToken, refreshToken, err := a.generateAllTokensConcurrent(ctx,
		acc.ID.String(),
		acc.Username,
		string(acc.Role))
		// createdActivity.ID.String())
	if err != nil {
		panic(err)
	}

	return token.Tokens{
		IDToken:      (idToken),
		AccessToken:  (accessToken),
		RefreshToken: (refreshToken),
	}, err
}

func (a *UserServiceImpl) GetAccountSvc(ctx context.Context, userId string) (account models.AccountResponse, err error) {
	// get account from database
	acc, err := a.userRepo.GetAccountByUserId(ctx, userId)
	if err != nil {
		return
	}
	return models.AccountResponse{
		ID:        acc.ID,
		Username:  acc.Username,
		Role:      acc.Role,
		CreatedAt: acc.CreatedAt,
	}, err
}
