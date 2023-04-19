package servers

import (
	// "fmt"
	// "log"
	"net/http"

	"github.com/Retrospective53/myGram/config"
	"github.com/Retrospective53/myGram/module/router/photo"
	"github.com/Retrospective53/myGram/module/router/user"
	"github.com/gin-gonic/gin"
)

func NewHttpServer() (srv *http.Server) {
	hdls := initDI()

	// init server
	ginServer := gin.Default()
	// if config.Load.Server.Env == config.ENV_PRODUCTION {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	// init middleware
	ginServer.Use(
		// gin.Logger(),                             // log request
		gin.Recovery(),                           // auto restart if panic
	)

	// register router
	v1 := ginServer.Group("/api/v1")
	user.NewAccountRouter(v1, hdls.accountHdl)
	photo.NewPhotoRouter(v1, hdls.photoHdl)
	ginServer.Run(config.Port)

	// srv = &http.Server{
	// 	Addr:    fmt.Sprintf(":%v", config.Load.Server.Http.Port),
	// 	Handler: ginServer,
	// }

	// go func() {
	// 	// service connections
	// 	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("listen: %s\n", err)
	// 	}
	// }()
	return
}