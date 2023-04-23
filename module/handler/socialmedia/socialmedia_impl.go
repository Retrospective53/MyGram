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


// @Tags Social Media
// @Summary finds all social media records
// @Schemes http
// @Description fetch all social media records
// @Param Authorization header string true "Bearer Token"
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]models.Socialmedia}
// @Success 401 {object} response.SuccessResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /socialmedias/all [get]
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


// @Tags Social Media
// @Summary Find a social media by ID
// @Schemes http
// @Description Fetch a social media with the given id
// @Accept json
// @Param id path string true "SocialMedia ID"
// @Param Authorization header string true "Bearer Token"
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=models.Socialmedia}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 401 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /socialmedias/{id} [get]
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

// @Tags Social Media
// @Summary Create a new social media
// @Schemes http
// @Description Creates a new social media with the provided data
// @Accept json
// @Param body body socialmediacreatemodel.SocialMediaCreate true "Create Social Media Request Body"
// @Param Authorization header string true "Bearer Token"
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=models.Socialmedia}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /socialmedias [post]
func (s *SocialMediaHandlerImpl) CreateSocialMediaHdl(ctx *gin.Context) {
	// get user_id from context first
	accessClaimI, ok := ctx.Get(middleware.AccessClaim.String())
	if !ok {
		err := errors.New("error get claim from context")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
				Message: response.SomethingWentWrong,
				Error:   err.Error(),
			})
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

// @Tags Social Media
// @Summary Update an existing social media by id
// @Schemes http
// @Description Updates an existing social media with the provided data
// @Accept json
// @Param id path string true "SocialMedia ID"
// @Param Authorization header string true "Bearer Token"
// @Param request body models.Comment true "Create Social Media Request Body"
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=models.Socialmedia}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 401 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /socialmedias/{id} [put]
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
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
				Message: response.SomethingWentWrong,
				Error:   err.Error(),
			})
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

	if socialMedia.Name == "" {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InvalidParam,
			Error:   "social media not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    socialMedia,
	})
}

// @Tags Social Media
// @Summary Delete a social media by ID
// @Schemes http
// @Description Deletes a social media with the given id
// @Accept json
// @Param id path string true "Social Media ID"
// @Param Authorization header string true "Bearer Token"
// @Produce json
// @Success 200 {object} response.SuccessResponse{}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 401 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /socialmedias/{id} [delete]
func (s *SocialMediaHandlerImpl) DeleteSocialMediaByIdHdl(ctx *gin.Context) {
	socialMediaId := ctx.Param("id")

	// get user_id from context first
	accessClaimI, ok := ctx.Get(middleware.AccessClaim.String())
	if !ok {
		err := errors.New("error get claim from context")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
				Message: response.SomethingWentWrong,
				Error:   err.Error(),
			})
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

