package comment

import (
	"context"

	"github.com/Retrospective53/myGram/module/models"
	commentCreateModel "github.com/Retrospective53/myGram/module/models/comment"
	commentrepo "github.com/Retrospective53/myGram/module/repository/comment"
	"github.com/google/uuid"
)

type CommentServiceImpl struct {
	commentRepo commentrepo.CommentRepo
}

func NewCommentServiceImpl(commentRepo commentrepo.CommentRepo) CommentService {
	return &CommentServiceImpl{
		commentRepo: commentRepo,
	}
}

func (c *CommentServiceImpl) FindAllCommentsSvc(ctx context.Context) (comments []models.Comment, err error) {
	comments, err = c.commentRepo.FindAllComment(ctx)
	if err != nil {
		panic(err)
	}

	return
}

func (c *CommentServiceImpl) FindCommentByIdSvc(ctx context.Context, commentId string) (comment models.Comment, err error) {
	comment, err = c.commentRepo.FindCommentById(ctx, commentId)
	if err != nil {
		panic(err)
	}

	return
}

func (c *CommentServiceImpl) CreateCommentSvc(ctx context.Context, commentIn commentCreateModel.CommentCreate, userId string) (comment models.Comment, err error) {
	// Convert userID from string to uuid.UUID
	userUUID, err := uuid.Parse(commentIn.UserID)
	if err != nil {
		return
	}

	photoUUID, err := uuid.Parse(commentIn.PhotoID)
	if err != nil {
		return
	}
	
	comment, err = c.commentRepo.CreateComment(ctx, models.Comment{
		ID: uuid.New(),
		Message: commentIn.Message,
		PhotoID: photoUUID,
		UserID: userUUID,
	}, userId)
	if err != nil {
		panic(err)
	}

	return
}

func (c *CommentServiceImpl) UpdateCommentSvc(ctx context.Context, commentIn models.Comment, commentId string) (comment models.Comment, err error) {
	comment, err = c.commentRepo.UpdateComment(ctx, commentIn, commentId)
	if err != nil {
		panic(err)
	}
	return
}

func (c *CommentServiceImpl) DeleteCommentByIdSvc(ctx context.Context, commentId string) (comment models.Comment, err error) {
	err = c.commentRepo.DeleteCommentById(ctx, commentId)
	if err != nil {
		panic(err)
	}
	return
}
