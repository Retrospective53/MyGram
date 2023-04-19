package servers

import (
	"github.com/Retrospective53/myGram/config"

	accounthdl "github.com/Retrospective53/myGram/module/handler/user"
	accountrepo "github.com/Retrospective53/myGram/module/repository/user"
	accountsvc "github.com/Retrospective53/myGram/module/service/user"

	photohdl "github.com/Retrospective53/myGram/module/handler/photo"
	photorepo "github.com/Retrospective53/myGram/module/repository/photo"
	photosvc "github.com/Retrospective53/myGram/module/service/photo"
	// c "github.com/Retrospective53/go-common/pkg/context"
)

type handlers struct {
	accountHdl accounthdl.UserHandler
	photoHdl photohdl.PhotoHandler
}

func initDI() handlers {
	pgConn := config.NewPostgresGormConn()
	accountRepo := accountrepo.NewAccountRepoGormImpl(pgConn)
	accountSvc := accountsvc.NewUserServiceImpl(accountRepo)
	accountHdl := accounthdl.NewAccountHandlerImpl(accountSvc)

	photoRepo := photorepo.NewPhotoRepoGormImpl(pgConn)
	photoSvc := photosvc.NewPhotoServiceImpl(photoRepo)
	photoHdl := photohdl.NewPhotoHandlerImpl(photoSvc)

	return handlers{
		accountHdl: accountHdl,
		photoHdl: photoHdl,
	}
}