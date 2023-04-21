package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Photo struct {
	ID        uuid.UUID      `json:"id" gorm:"column:id;type:char(36);primaryKey"`
	Title     string         `json:"title" gorm:"column:title;not null" valid:"required~Title is required"`
	Caption   string         `json:"caption" gorm:"column:caption"`
	PhotoURL  string         `json:"photo_url" gorm:"column:photo_url;not null" valid:"required~Photo url is required"`
	UserID    uuid.UUID      `json:"user_id" gorm:"column:user_id;type:char(36)"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
	Comments  []Comment      `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}