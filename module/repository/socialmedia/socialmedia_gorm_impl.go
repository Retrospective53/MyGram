package socialmedia

import (
	"context"

	"github.com/Retrospective53/myGram/module/models"
	"gorm.io/gorm"
)

type SocialMediaRepoGormImpl struct {
	master *gorm.DB
}

func NewSocialMediaRepoGormImpl(master *gorm.DB) SocialMediaRepo {
	return &SocialMediaRepoGormImpl{
		master: master,
	}
}

func(s *SocialMediaRepoGormImpl) FindAllSocialMedia(ctx context.Context) (socialMedias []models.Socialmedia, err error) {
	err = s.master.
		Table("socialmedia").
		Find(&socialMedias).
		Order("id ASC").
		Error

	return
}

func(s *SocialMediaRepoGormImpl) FindSocialMediaById(ctx context.Context, socialMediaId string) (socialMedia models.Socialmedia, err error) {
	err = s.master.
		Table("socialmedia").
		Where("id = ?", socialMediaId).
		Find(&socialMedia).
		Error

	return
}

func(s *SocialMediaRepoGormImpl) CreateSocialMedia(ctx context.Context, socialMediaIn models.Socialmedia, socialMediaId string) (socialMedia models.Socialmedia, err error) {
	err = s.master.
		Table("socialmedia").
		Create(&socialMediaIn).
		Error

	return socialMediaIn, err
}

func(s *SocialMediaRepoGormImpl) UpdateSocialMedia(ctx context.Context, socialMediaIn models.Socialmedia, socialMediaId string) (socialMedia models.Socialmedia, err error) {
	err = s.master.
		Table("socialmedia").
		Where("id = ?", socialMediaId).
		Updates(&socialMediaIn).
		Find(&socialMedia).
		Error
	
	return
}

func(s *SocialMediaRepoGormImpl) DeleteSocialMediaById(ctx context.Context, socialMediaId string) (err error) {
	err = s.master.
		Table("socialmedia").
		Where("id = ?", socialMediaId).
		Delete(&models.Socialmedia{}).
		Error

	return
}