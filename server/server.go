package server

import (
	"github.com/IQ-tech/go-mapper"
	"github.com/gin-gonic/gin"
	"github.com/safra-team-35/backend/server/routes/pingroute"
	"github.com/safra-team-35/backend/server/routes/qrcoderoute"
	"github.com/safra-team-35/backend/service"
)

type controller struct {
	pingController   *pingroute.Controller
	qrcodeController *qrcoderoute.Controller
}

//InitServer to initialize the server
func InitServer(svc *service.Service) *gin.Engine {
	mapper := mapper.New()
	svm := service.NewServiceManager()
	srv := gin.Default()

	qrcodeService := svm.QRCodeService(svc)

	return setupRoutes(srv, &controller{
		pingController:   pingroute.NewController(),
		qrcodeController: qrcoderoute.NewController(qrcodeService, mapper),
	})
}

//setupRoutes - Register and instantiate the routes
func setupRoutes(srv *gin.Engine, s *controller) *gin.Engine {

	pingroute.NewRouter(s.pingController, srv).RegisterRoutes()
	qrcoderoute.NewRouter(s.qrcodeController, srv).RegisterRoutes()

	return srv
}
