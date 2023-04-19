package photo

import (
	"context"

	"github.com/Retrospective53/myGram/module/models"
	photoCreateModel "github.com/Retrospective53/myGram/module/models/photo"
)

type PhotoService interface {
	FindAllPhotosSvc(ctx context.Context) (photos []models.Photo, err error)
	FindPhotoByIdSvc(ctx context.Context, photoId string) (photo models.Photo, err error)
	CreatePhotoSvc(ctx context.Context, photoIn photoCreateModel.PhotoCreate, userId string) (photo models.Photo, err error)
	UpdatePhotoSvc(ctx context.Context, photoIn models.Photo, photoId string) (photo models.Photo, err error)
	DeletePhotoByIdSvc(ctx context.Context, photoId string) (photo models.Photo, err error)
}