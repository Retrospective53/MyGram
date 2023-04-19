package comment

import (
	"context"

	"github.com/Retrospective53/myGram/module/models"
	"gorm.io/gorm"
)

type CommentRepoGormImpl struct {
	master *gorm.DB
}

func NewCommentRepoGormImpl(master *gorm.DB) CommentRepo {
	return &CommentRepoGormImpl{
		master: master,
	}
}

func (c *CommentRepoGormImpl) FindAllComment(ctx context.Context) (comments []models.Comment, err error) {
	err = c.master.
		Table("comments").
		Find(&comments).
		Order("id ASC").
		Error
		
	return
}

func (c *CommentRepoGormImpl) FindCommentById(ctx context.Context, commentId string) (comment models.Comment, err error) {
	err = c.master.
		Table("comments").
		Where("id = ?", commentId).
		Find(&comment).
		Error

	return
}

func (c *CommentRepoGormImpl) CreateComment(ctx context.Context, commentIn models.Comment, commentId string) (comment models.Comment, err error) {
	err = c.master.
		Table("comments").
		Create(&commentIn).
		Error

	return commentIn, err
	}

func (c *CommentRepoGormImpl) UpdateComment(ctx context.Context, commentIn models.Comment, commentId string) (comment models.Comment, err error)  {
	err = c.master.
		Table("comments").
		Where("id = ?", commentId).
		Updates(&commentIn).
		Find(&comment).
		Error

	return
}

func (c *CommentRepoGormImpl) DeleteCommentById(ctx context.Context, commentId string) (err error) {
	err = c.master.
		Table("comments").
		Where("id = ?", commentId).
		Delete(&models.Comment{}).
		Error
	
	return	
}