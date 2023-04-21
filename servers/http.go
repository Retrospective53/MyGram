package servers

import (
	// "fmt"
	// "log"
	"net/http"

	"github.com/Retrospective53/myGram/config"
	"github.com/Retrospective53/myGram/module/router/comment"
	"github.com/Retrospective53/myGram/module/router/photo"
	"github.com/Retrospective53/myGram/module/router/socialmedia"
	"github.com/Retrospective53/myGram/module/router/user"
	"github.com/gin-gonic/gin"

	docs "github.com/Retrospective53/myGram/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyGram API DUCUMENTATION
// @version 1.0
// @description mygram api documentation
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @schemes http
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
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := ginServer.Group("/api/v1")

	// swagger
	ginServer.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user.NewAccountRouter(v1, hdls.accountHdl)
	photo.NewPhotoRouter(v1, hdls.photoHdl)
	comment.NewCommentRouter(v1, hdls.commentHdl)
	socialmedia.NewSocialMediaRouter(v1, hdls.socialMediaHdl)
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