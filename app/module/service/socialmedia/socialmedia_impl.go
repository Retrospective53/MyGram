package socialmedia

import (
	"context"

	"github.com/Retrospective53/myGram/module/models"
	socialMediaCreateModel "github.com/Retrospective53/myGram/module/models/socialmedia"
	socialmediarepo "github.com/Retrospective53/myGram/module/repository/socialmedia"
	"github.com/google/uuid"
)

type SocialMediaServiceImpl struct {
	socialMediaRepo socialmediarepo.SocialMediaRepo
}

func NewSocialMediaServiceImpl(socialMediaRepo socialmediarepo.SocialMediaRepo) SocialMediaService {
	return &SocialMediaServiceImpl{
		socialMediaRepo: socialMediaRepo,
	}
}

func (s *SocialMediaServiceImpl) FindAllSocialMediaSvc(ctx context.Context) (socialMedias []models.Socialmedia, err error) {
	socialMedias, err = s.socialMediaRepo.FindAllSocialMedia(ctx)
	if err != nil {
		panic(err)
	}

	return
}

func (s *SocialMediaServiceImpl) FindSocialMediaByIdSvc(ctx context.Context, socialMediaId string) (socialMedia models.Socialmedia, err error) {
	socialMedia, err = s.socialMediaRepo.FindSocialMediaById(ctx, socialMediaId)
	if err != nil {
		panic(err)
	}

	return
}

func (s *SocialMediaServiceImpl) CreateSocialMediaSvc(ctx context.Context, socialMediaIn socialMediaCreateModel.SocialMediaCreate) (socialMedia models.Socialmedia, err error) {
		// Convert userID from string to uuid.UUID
		userUUID, err := uuid.Parse(socialMediaIn.UserID)
		if err != nil {
			return
		}
	
		socialMedia, err = s.socialMediaRepo.CreateSocialMedia(ctx, models.Socialmedia{
			ID: uuid.New(),
			Name: socialMediaIn.Name,
			SocialMediaURL: socialMediaIn.SocialMediaURL,
			UserID: userUUID,
		}, socialMediaIn.UserID)

		if err != nil {
			panic(err)
		}
	
	return
}

func (s *SocialMediaServiceImpl) UpdateSocialMediaSvc(ctx context.Context, socialMediaIn models.Socialmedia, socialMediaId string) (socialMedia models.Socialmedia, err error) {
	socialMedia, err = s.socialMediaRepo.UpdateSocialMedia(ctx, socialMediaIn, socialMediaId)
	if err != nil {
		panic(err)
	}
	
	return
}

func (s *SocialMediaServiceImpl) DeleteSocialMediaByIdSvc(ctx context.Context, socialMediaId string) (socialMedia models.Socialmedia, err error) {
	err = s.socialMediaRepo.DeleteSocialMediaById(ctx, socialMediaId)
	if err != nil {
		panic(err)
	}
	
	return
}
