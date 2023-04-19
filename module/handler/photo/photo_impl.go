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

	photo, err := p.photoService.UpdatePhotoSvc(ctx, updatePhoto, photoId)
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

func (p *PhotoHandlerImpl) DeletePhotoByIdHdl(ctx *gin.Context) {
	photoId := p.getIdFromParamStr(ctx)

	photo, err := p.photoService.DeletePhotoByIdSvc(ctx, photoId)
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