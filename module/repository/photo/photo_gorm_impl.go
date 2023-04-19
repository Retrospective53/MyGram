package photo

import (
	"context"

	"github.com/Retrospective53/myGram/module/models"
	"gorm.io/gorm"
)

type PhotoRepoGormImpl struct {
	master *gorm.DB
}

func NewPhotoRepoGormImpl(master *gorm.DB) PhotoRepo {
	return &PhotoRepoGormImpl{
		master: master,
	}
}


func (p *PhotoRepoGormImpl) FindAllPhotos(ctx context.Context) (photos []models.Photo, err error) {
	err = p.master.
		Table("photos").
		Preload("Comments").
		Find(&photos).
		Order("id ASC").
		Error
		
	return
}

func (p *PhotoRepoGormImpl) FindPhotoById(ctx context.Context, photoId string) (photo models.Photo, err error) {
	err = p.master.
		Table("photos").
		Where("id = ?", photoId).
		Preload("Comments").
		Find(&photo).
		Error

		return
}

func (p *PhotoRepoGormImpl) CreatePhoto(ctx context.Context, photoIn models.Photo, userId string) (photo models.Photo, err error) {
	err = p.master.
		Table("photos").
		Create(&photoIn).
		Error

	return photoIn, err
}

func (p *PhotoRepoGormImpl) UpdatePhoto(ctx context.Context, photoIn models.Photo, photoId string) (photo models.Photo, err error) {
	err = p.master.
		Table("photos").
		Where("id = ?", photoId).
		Updates(&photoIn).
		Find(&photo).
		Error

	return
}

func (p *PhotoRepoGormImpl) DeletePhotoById(ctx context.Context, photoId string) (err error) {
	err = p.master.
		Table("photos").
		Where("id = ?", photoId).
		Delete(&models.Photo{}).
		Error
	
	return
}
