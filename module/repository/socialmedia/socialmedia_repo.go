package socialmedia

import (
	"context"

	"github.com/Retrospective53/myGram/module/models"
)

type SocialMediaRepo interface {
	FindAllSocialMedia(ctx context.Context) (socialMedias []models.Socialmedia, err error)
	FindSocialMediaById(ctx context.Context, socialMediaId string) (socialMedia models.Socialmedia, err error)
	CreateSocialMedia(ctx context.Context, socialMediaIn models.Socialmedia, socialMediaId string) (socialMedia models.Socialmedia, err error)
	UpdateSocialMedia(ctx context.Context, socialMediaIn models.Socialmedia, socialMediaId string) (socialMedia models.Socialmedia, err error)
	DeleteSocialMediaById(ctx context.Context, socialMediaId string) (err error)
}