package user

import (
	"context"
	"log"

	"github.com/Retrospective53/myGram/config"
	"github.com/Retrospective53/myGram/module/models"
	"gorm.io/gorm"
)

type AccountRepoGormImpl struct {
	master *gorm.DB
}

func NewAccountRepoGormImpl(master *gorm.DB) UserRepo {
	userRepo := AccountRepoGormImpl{
		master: master,
	}

	if config.MIGRATE {
		userRepo.doMigration()
	}

	return &AccountRepoGormImpl{
		master: master,
	}
}

func (a *AccountRepoGormImpl) doMigration() (err error) {
	if err = a.master.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.Socialmedia{}); err != nil {
		panic(err)
	}

	log.Println("succesfully create book table")
	return
}

func (a *AccountRepoGormImpl) CreateAccount(ctx context.Context, acc models.User) (created models.User, err error) {
	err = a.master.
		Table("users").
		Create(&acc).Error
	if err != nil {
		return
	}

	return acc, err
}

func (a *AccountRepoGormImpl) GetAccountByUserName(ctx context.Context, username string) (account models.User, err error) {
	err = a.master.
		Table("users").
		Where("username = ?", username).
		Preload("Photos").
		Find(&account).Error
	if err != nil {
		return
	}

	return account, err
}

func (a *AccountRepoGormImpl) GetAccountByUserId(ctx context.Context, userId string) (account models.User, err error) {
	err = a.master.
		Table("users").
		Where("id = ?", userId).
		Preload("Photos").
		Find(&account).Error

	if err != nil {
		return
	}
	return account, err
}