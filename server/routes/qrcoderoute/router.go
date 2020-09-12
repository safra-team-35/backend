package qrcoderoute

import (
	"github.com/gin-gonic/gin"
)

// QRCodeRouter holds the QRCode handlers
type QRCodeRouter struct {
	ctrl   *Controller
	router *gin.Engine
}

// NewRouter returns a new PingRouter instance
func NewRouter(ctrl *Controller, router *gin.Engine) *QRCodeRouter {
	return &QRCodeRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of qrcode requests
func (r *QRCodeRouter) RegisterRoutes() {
	r.router.POST("/partner/:company_id/qrcode", r.ctrl.handleCreateQRCode)
	r.router.GET("/qrcode/:hash", r.ctrl.handleGetQRCodeDataByHash)
}
