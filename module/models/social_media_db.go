package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Socialmedia struct {
	ID         uuid.UUID      `json:"id" gorm:"column:id;type:char(36);primaryKey"`
	Name       string         `json:"name" gorm:"column:name;not null" valid:"required~Socialmedia name is required"`
	SocialMediaURL string    `json:"social_media_url" gorm:"column:social_media_url;not null" valid:"required~Social media url is required"`
	UserID     uuid.UUID      `json:"user_id"`
	User      *User           `json:"user"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

func (p *Socialmedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Socialmedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}