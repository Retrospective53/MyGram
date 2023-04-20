package socialmedia

import (
	"errors"
	"log"
	"net/http"

	"github.com/Retrospective53/myGram/module/models"
	socialmediacreatemodel "github.com/Retrospective53/myGram/module/models/socialmedia"
	"github.com/Retrospective53/myGram/module/models/token"
	socialmeidaservice "github.com/Retrospective53/myGram/module/service/socialmedia"
	"github.com/Retrospective53/myGram/pkg/json"
	"github.com/Retrospective53/myGram/pkg/middleware"
	"github.com/Retrospective53/myGram/pkg/response"
	"github.com/gin-gonic/gin"
)

type SocialMediaHandlerImpl struct {
	socialMediaService socialmeidaservice.SocialMediaService
}

func NewSocialMediaHandlerImpl(socialMediaService socialmeidaservice.SocialMediaService) SocialMediaHandler {
	return &SocialMediaHandlerImpl{
		socialMediaService: socialMediaService,
	}
}

func (s *SocialMediaHandlerImpl) FindAllSocialMediasHdl(ctx *gin.Context) {
	socialMedias, err := s.socialMediaService.FindAllSocialMediaSvc(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InternalServer,
			Error:   "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    socialMedias,
	})
}

func (s *SocialMediaHandlerImpl) FindSocialMediaByIdHdl(ctx *gin.Context) {
	socialMediaId := ctx.Param("id")

	socialMedia, err := s.socialMediaService.FindSocialMediaByIdSvc(ctx, socialMediaId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InternalServer,
			Error:   "something went wrong",
		})
		return
	}

	if socialMedia.Name == "" || socialMedia.ID.String() == "" {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InvalidParam,
			Error:   "photo not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    socialMedia,
	})
}

func (s *SocialMediaHandlerImpl) CreateSocialMediaHdl(ctx *gin.Context) {
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

		// binding payload
		var createSocialMedia socialmediacreatemodel.SocialMediaCreate
		createSocialMedia.UserID = accessClaim.UserID
		log.Printf("%s data type is: %T", accessClaim.UserID, accessClaim.UserID)
		if err := ctx.BindJSON(&createSocialMedia); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				response.ErrorResponse{
					Message: response.InvalidBody,
					Error:   "error binding payload",
				},
			)
			return
		}

	socialMedia, err := s.socialMediaService.CreateSocialMediaSvc(ctx, createSocialMedia)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InternalServer,
			Error:   "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    socialMedia,
	})
}

func (s *SocialMediaHandlerImpl) UpdateSocialMediaHdl(ctx *gin.Context) {
	socialMediaId := ctx.Param("id")
	
	// binding payload
	var updateSocialMedia models.Socialmedia
	if err := ctx.BindJSON(&updateSocialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			response.ErrorResponse{
				Message: response.InvalidBody,
				Error:   "error binding payload",
			},
		)
		return
	}

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


	// authorization only admin
	if accessClaim.Role != "ROLE_ADMIN" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
			Message: response.Unauthorized,
			Error: "Unauthorized",
		})
		return
	}

	socialMedia, err := s.socialMediaService.UpdateSocialMediaSvc(ctx, updateSocialMedia, socialMediaId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InternalServer,
			Error:   "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    socialMedia,
	})
}

func (s *SocialMediaHandlerImpl) DeleteSocialMediaByIdHdl(ctx *gin.Context) {
	socialMediaId := ctx.Param("id")

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


	// authorization only admin
	if accessClaim.Role != "ROLE_ADMIN" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
			Message: response.Unauthorized,
			Error: "Unauthorized",
		})
		return
	}

	_, err := s.socialMediaService.DeleteSocialMediaByIdSvc(ctx, socialMediaId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InternalServer,
			Error:   "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    "social media deleted",
	})
}

