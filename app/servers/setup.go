package servers

import (
	"github.com/Retrospective53/myGram/config"

	accounthdl "github.com/Retrospective53/myGram/module/handler/user"
	accountrepo "github.com/Retrospective53/myGram/module/repository/user"
	accountsvc "github.com/Retrospective53/myGram/module/service/user"

	photohdl "github.com/Retrospective53/myGram/module/handler/photo"
	photorepo "github.com/Retrospective53/myGram/module/repository/photo"
	photosvc "github.com/Retrospective53/myGram/module/service/photo"

	commenthdl "github.com/Retrospective53/myGram/module/handler/comment"
	commentrepo "github.com/Retrospective53/myGram/module/repository/comment"
	commentsvc "github.com/Retrospective53/myGram/module/service/comment"

	socialmediahdl "github.com/Retrospective53/myGram/module/handler/socialmedia"
	socialmediarepo "github.com/Retrospective53/myGram/module/repository/socialmedia"
	socialmediasvc "github.com/Retrospective53/myGram/module/service/socialmedia"
	// c "github.com/Retrospective53/go-common/pkg/context"
)

type handlers struct {
	accountHdl accounthdl.UserHandler
	photoHdl photohdl.PhotoHandler
	commentHdl commenthdl.CommentHandler
	socialMediaHdl socialmediahdl.SocialMediaHandler
}

func initDI() handlers {
	pgConn := config.NewPostgresGormConn()
	accountRepo := accountrepo.NewAccountRepoGormImpl(pgConn)
	accountSvc := accountsvc.NewUserServiceImpl(accountRepo)
	accountHdl := accounthdl.NewAccountHandlerImpl(accountSvc)

	photoRepo := photorepo.NewPhotoRepoGormImpl(pgConn)
	photoSvc := photosvc.NewPhotoServiceImpl(photoRepo)
	photoHdl := photohdl.NewPhotoHandlerImpl(photoSvc)

	commentRepo := commentrepo.NewCommentRepoGormImpl(pgConn)
	commentSvc := commentsvc.NewCommentServiceImpl(commentRepo)
	commentHdl := commenthdl.NewCommentHandlerImpl(commentSvc)

	socialMediaRepo := socialmediarepo.NewSocialMediaRepoGormImpl(pgConn)
	socialMediaSvc := socialmediasvc.NewSocialMediaServiceImpl(socialMediaRepo)
	socialMediaHdl := socialmediahdl.NewSocialMediaHandlerImpl(socialMediaSvc)

	return handlers{
		accountHdl: accountHdl,
		photoHdl: photoHdl,
		commentHdl: commentHdl,
		socialMediaHdl: socialMediaHdl,
	}
}