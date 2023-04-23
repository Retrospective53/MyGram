package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRole string

const (
	ROLE_ADMIN  AccountRole = "admin"
	ROLE_NORMAL AccountRole = "normal"
)

type User struct {
	ID        uuid.UUID      `json:"id" gorm:"column:id;type:char(36);primaryKey"`
	Username  string         `json:"username" gorm:"column:username;not null;uniqueIndex" valid:"required~Username is required"`
	Email     string         `json:"email" gorm:"column:email;not null;uniqueIndex" valid:"required~Your email is required,email~Invalid email format"`
	Password  string         `json:"password" gorm:"column:password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age       int            `json:"age" gorm:"column:age;not null" valid:"required~Age is required,range(8|200)~Age must be at least 8 years old"`
	Role      AccountRole    `json:"role" gorm:"column:role"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
	Photos    []Photo        `json:"photos" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Comments  []Comment      `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Socialmedias []Socialmedia `json:"socialmedias" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (p *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}