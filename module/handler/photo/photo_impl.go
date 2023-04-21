package photo

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Retrospective53/myGram/module/models"
	photoCreateModel "github.com/Retrospective53/myGram/module/models/photo"
	"github.com/Retrospective53/myGram/module/models/token"
	photoservice "github.com/Retrospective53/myGram/module/service/photo"
	"github.com/Retrospective53/myGram/pkg/json"
	"github.com/Retrospective53/myGram/pkg/middleware"
	"github.com/Retrospective53/myGram/pkg/response"
	"github.com/gin-gonic/gin"
)

type PhotoHandlerImpl struct {
	photoService photoservice.PhotoService
}

func NewPhotoHandlerImpl(photoService photoservice.PhotoService) PhotoHandler {
	return &PhotoHandlerImpl{
		photoService: photoService,
	}
}



// @BasePath /api/v1/photo

// @Tags Photo
// @Summary finds all photo records
// @Schemes http
// @Description fetch all photo records
// @Accept json
// @Param Authorization header string true "Bearer Token"
// @Produce json
// @Success 200 {object} response.SuccessResponse{}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /photo/all [get]

// @Tags Photo
// @Summary finds all photo records
// @Schemes http
// @Description fetch all photo records
// @Param Authorization header string true "Bearer Token"
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]string}
// @Success 401 {object} response.SuccessResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /photo/all [get]
func (p *PhotoHandlerImpl) FindAllPhotosHdl(ctx *gin.Context) {
	photos, err := p.photoService.FindAllPhotosSvc(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InternalServer,
			Error:   "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    photos,
	})
}

// @Tags Photo
// @Summary Find a photo by ID
// @Schemes http
// @Description Fetch a photo with the given id
// @Accept json
// @Param id path string true "Photo ID"
// @Param Authorization header string true "Bearer Token"
// @Produce json
// @Success 200 {object} response.SuccessResponse{}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 401 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /photo/{id} [get]
func (p *PhotoHandlerImpl) FindPhotoByIdHdl(ctx *gin.Context) {
	photoId := p.getIdFromParamStr(ctx)

	photo, err := p.photoService.FindPhotoByIdSvc(ctx, photoId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InternalServer,
			Error:   "something went wrong",
		})
		return
	}

	if photo.Title == "" || photo.PhotoURL == "" {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InvalidParam,
			Error:   "photo not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    photo,
	})
}


// @Tags Photo
// @Summary Create a new photo
// @Schemes http
// @Description Creates a new photo with the provided data
// @Accept json
// @Param body body photoCreateModel.PhotoCreate true "Create Photo Request Body"
// @Param Authorization header string true "Bearer Token"
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=object}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /photo [post]
func (p *PhotoHandlerImpl) CreatePhotoHdl(ctx *gin.Context) {
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
		var createPhoto photoCreateModel.PhotoCreate
		createPhoto.UserID = accessClaim.UserID
		log.Printf("%s data type is: %T", accessClaim.UserID, accessClaim.UserID)
		if err := ctx.BindJSON(&createPhoto); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				response.ErrorResponse{
					Message: response.InvalidBody,
					Error:   "error binding payload",
				},
			)
			return
		}

	photo, err := p.photoService.CreatePhotoSvc(ctx, createPhoto, accessClaim.UserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InternalServer,
			Error:   "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    photo,
	})
}


// @Tags Photo
// @Summary Update an existing photo by id
// @Schemes http
// @Description Updates an existing photo with the provided data
// @Accept json
// @Param id path string true "Photo ID"
// @Param Authorization header string true "Bearer Token"
// @Param request body models.Photo true "Create Photo Request Body"
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=object}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 401 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /photo/{id} [put]
func (p *PhotoHandlerImpl) UpdatePhotoHdl(ctx *gin.Context) {
	photoId := p.getIdFromParamStr(ctx)
	
	// binding payload
	var updatePhoto models.Photo
	if err := ctx.BindJSON(&updatePhoto); err != nil {
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
			Error: "Unauthorized only admin is allowed lol",
		})
		return
	}

	photo, err := p.photoService.UpdatePhotoSvc(ctx, updatePhoto, photoId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InternalServer,
			Error:   "something went wrong",
		})
		return
	}

	if photo.PhotoURL == "" {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InvalidParam,
			Error:   "photo not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    photo,
	})
}


// @Tags Photo
// @Summary Delete a photo by ID
// @Schemes http
// @Description Deletes a photo with the given id
// @Accept json
// @Param id path string true "Photo ID"
// @Param Authorization header string true "Bearer Token"
// @Produce json
// @Success 200 {object} response.SuccessResponse{}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 401 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /photo/{id} [delete]
func (p *PhotoHandlerImpl) DeletePhotoByIdHdl(ctx *gin.Context) {
	photoId := p.getIdFromParamStr(ctx)

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

	_, err := p.photoService.DeletePhotoByIdSvc(ctx, photoId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InternalServer,
			Error:   "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    "photo deleted",
	})
}

func (p *PhotoHandlerImpl) getIdFromParam(ctx *gin.Context) (idUint uint64, err error) {
	id := ctx.Param("id")

	// transform id string to uint64
	idUint, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to find photo",
			Error: response.InvalidParam,
		})
		return
	}

	return
}

func (p *PhotoHandlerImpl) getIdFromParamStr(ctx *gin.Context) (id string) {
	id = ctx.Param("id")

	// // transform id string to uint64
	// idUint, err = strconv.ParseUint(id, 10, 64)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
	// 		Message:  "failed to find photo",
	// 		Error: response.InvalidParam,
	// 	})
	// 	return
	// }

	return
}