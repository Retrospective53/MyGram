package socialmedia

import (
	"context"

	"github.com/Retrospective53/myGram/module/models"
	socialMediaCreateModel "github.com/Retrospective53/myGram/module/models/socialmedia"
)

type SocialMediaService interface {
	FindAllSocialMediaSvc(ctx context.Context) (socialMedias []models.Socialmedia, err error)
	FindSocialMediaByIdSvc(ctx context.Context, socialMediaId string) (socialMedia models.Socialmedia, err error)
	CreateSocialMediaSvc(ctx context.Context, socialMediaIn socialMediaCreateModel.SocialMediaCreate) (socialMedia models.Socialmedia, err error)
	UpdateSocialMediaSvc(ctx context.Context, socialMediaIn models.Socialmedia, socialMediaId string) (socialMedia models.Socialmedia, err error)
	DeleteSocialMediaByIdSvc(ctx context.Context, socialMediaId string) (socialMedia models.Socialmedia, err error)
}