package photo

import (
	"context"

	"github.com/Retrospective53/myGram/module/models"
)

type PhotoRepo interface {
	FindAllPhotos(ctx context.Context) (photos []models.Photo, err error)
	FindPhotoById(ctx context.Context, photoId string) (photo models.Photo, err error)
	CreatePhoto(ctx context.Context, photoIn models.Photo, userId string) (photo models.Photo, err error)
	UpdatePhoto(ctx context.Context, photoIn models.Photo, photoId string) (photo models.Photo, err error)
	DeletePhotoById(ctx context.Context, photoId string) (err error)
}