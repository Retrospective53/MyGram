package photo

import (
	"context"

	"github.com/Retrospective53/myGram/module/models"
	photoCreateModel "github.com/Retrospective53/myGram/module/models/photo"
	photorepo "github.com/Retrospective53/myGram/module/repository/photo"
	"github.com/google/uuid"
)

type PhotoServiceImpl struct {
	photoRepo photorepo.PhotoRepo
}

func NewPhotoServiceImpl(photoRepo photorepo.PhotoRepo) PhotoService {
	return &PhotoServiceImpl{
		photoRepo: photoRepo,
	}
}

func (p *PhotoServiceImpl) FindAllPhotosSvc(ctx context.Context) (photos []models.Photo, err error) {
	photos, err = p.photoRepo.FindAllPhotos(ctx)
	if err != nil {
		panic(err)
	}

	return
}

func (p *PhotoServiceImpl) FindPhotoByIdSvc(ctx context.Context, photoId string) (photo models.Photo, err error) {
	photo, err = p.photoRepo.FindPhotoById(ctx, photoId)
	if err != nil {
		panic(err)
	}

	return
}

func (p *PhotoServiceImpl) CreatePhotoSvc(ctx context.Context, photoIn photoCreateModel.PhotoCreate, userId string) (photo models.Photo, err error) {
	// Convert userID from string to uuid.UUID
	userUUID, err := uuid.Parse(photoIn.UserID)
	if err != nil {
		return
	}
	
	photo, err = p.photoRepo.CreatePhoto(ctx, models.Photo{
		ID: uuid.New(),
		Title: photoIn.Title,
		Caption: photoIn.Caption,
		PhotoURL: photoIn.PhotoURL,
		UserID: userUUID,
	}, userId)
	if err != nil {
		panic(err)
	}

	return
}

func (p *PhotoServiceImpl) UpdatePhotoSvc(ctx context.Context, photoIn models.Photo, photoId string) (photo models.Photo, err error) {
	photo, err = p.photoRepo.UpdatePhoto(ctx, photoIn, photoId)
	if err != nil {
		panic(err)
	}
	return
}

func (p *PhotoServiceImpl) DeletePhotoByIdSvc(ctx context.Context, photoId string) (photo models.Photo, err error) {
	err = p.photoRepo.DeletePhotoById(ctx, photoId)
	if err != nil {
		panic(err)
	}
	return
}
