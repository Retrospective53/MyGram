package user

import (
	"errors"
	"net/http"

	"github.com/Retrospective53/myGram/module/models"
	"github.com/Retrospective53/myGram/module/models/token"
	accountservice "github.com/Retrospective53/myGram/module/service/user"
	"github.com/Retrospective53/myGram/pkg/json"
	"github.com/Retrospective53/myGram/pkg/middleware"
	"github.com/Retrospective53/myGram/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserHandlerImpl struct {
	accService accountservice.UserService
}

func NewAccountHandlerImpl(accService accountservice.UserService) UserHandler {
	return &UserHandlerImpl{
		accService: accService,
	}
}



// @BasePath /api/v1/account

// @Tags User
// @Summary finding user record
// @Schemes http
// @Description fetch user information by id
// @Accept json
// @Param body body models.LoginAccount true "Login Account Request Body"
// @Produce json
// @Success 202 {object} response.SuccessResponse{data=[]string}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /account/login [post]
func (a *UserHandlerImpl) LoginAccount(ctx *gin.Context) {
	// binding payload
	var loginAccount models.LoginAccount
	if err := ctx.BindJSON(&loginAccount); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			response.ErrorResponse{
				Message: response.InvalidBody,
				Error:   "error binding payload",
			},
		)
		return
	}
	tokens, err := a.accService.LoginAccountByUserNameSvc(ctx, loginAccount)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			response.ErrorResponse{
				Message: response.InternalServer,
				Error:   response.SomethingWentWrong,
			},
		)
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Message: "login success",
		Data:    tokens,
	})
}



// @Tags User
// @Summary create user
// @Schemes http
// @Description create user by inputing the correct user datas
// @Accept json
// @Param body body models.CreateAccount true "Create Account Request Body"
// @Produce json
// @Success 202 {object} response.SuccessResponse{data=object}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /account [post]
func (a *UserHandlerImpl) CreateAccount(ctx *gin.Context) {
	// binding payload
	var createAccount models.CreateAccount
	if err := ctx.BindJSON(&createAccount); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			response.ErrorResponse{
				Message: response.InvalidBody,
				Error:   "error binding payload",
			},
		)
		return
	}

	created, err := a.accService.CreateAccountSvc(ctx, createAccount)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			response.ErrorResponse{
				Message: response.InternalServer,
				Error:   response.SomethingWentWrong,
			},
		)
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Message: "account successfully created",
		Data:    created,
	})
}


// @Tags User
// @Summary get user account 
// @Schemes http
// @Description get an user account
// @Param Authorization header string true "Bearer Token"
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=object}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 401 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /account [get]
func (a *UserHandlerImpl) GetAccount(ctx *gin.Context) {
	// get user_id from context first
	accessClaimI, ok := ctx.Get(middleware.AccessClaim.String())
	if !ok {
		err := errors.New("error get claim from context")
		if err != nil {
			panic(err)
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: response.InvalidPayload,
			Error:   "invalid user id",
		})
		return
	}

	var accessClaim token.AccessClaim
	if err := json.ObjectMapper(accessClaimI, &accessClaim); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: response.InvalidPayload,
			Error:   "invalid payload",
		})
		return
	}

	account, err := a.accService.GetAccountSvc(ctx, accessClaim.UserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InternalServer,
			Error:   "something went wrong",
		})
		return
	}

	if account.Username == "" || account.Role == "" {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InvalidParam,
			Error:   "account not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    account,
	})
}