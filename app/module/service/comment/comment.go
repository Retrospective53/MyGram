package comment

import (
	"context"

	"github.com/Retrospective53/myGram/module/models"
	commentCreateModel "github.com/Retrospective53/myGram/module/models/comment"
)

type CommentService interface {
	FindAllCommentsSvc(ctx context.Context) (comments []models.Comment, err error)
	FindCommentByIdSvc(ctx context.Context, commentId string) (comment models.Comment, err error)
	CreateCommentSvc(ctx context.Context, commentIn commentCreateModel.CommentCreate, userId string) (comment models.Comment, err error)
	UpdateCommentSvc(ctx context.Context, commentIn models.Comment, commentId string) (comment models.Comment, err error)
	DeleteCommentByIdSvc(ctx context.Context, commentId string) (comment models.Comment, err error)
}