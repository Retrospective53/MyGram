package comment

import (
	"context"

	"github.com/Retrospective53/myGram/module/models"
)

type CommentRepo interface {
	FindAllComment(ctx context.Context) (comments []models.Comment, err error)
	FindCommentById(ctx context.Context, commentId string) (comment models.Comment, err error)
	CreateComment(ctx context.Context, commentIn models.Comment, commentId string) (comment models.Comment, err error)
	UpdateComment(ctx context.Context, commentIn models.Comment, commentId string) (comment models.Comment, err error)
	DeleteCommentById(ctx context.Context, commentId string) (err error)
}