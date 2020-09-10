package server

import (
	"github.com/gin-gonic/gin"
	"github.com/safra-team-35/backend/server/routes/pingroute"
	"github.com/safra-team-35/backend/service"
)

type controller struct {
	pingController *pingroute.Controller
}

//InitServer to initialize the server
func InitServer(svc *service.Service) *gin.Engine {
	//svm := service.NewServiceManager()
	srv := gin.Default()

	return setupRoutes(srv, &controller{
		pingController: pingroute.NewController(),
	})
}

//setupRoutes - Register and instantiate the routes
func setupRoutes(srv *gin.Engine, s *controller) *gin.Engine {

	pingroute.NewRouter(s.pingController, srv).RegisterRoutes()

	return srv
}
