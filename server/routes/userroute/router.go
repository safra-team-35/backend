package userroute

import (
	"github.com/gin-gonic/gin"
)

// UserRouter holds the user handlers
type UserRouter struct {
	ctrl   *Controller
	router *gin.Engine
}

// NewRouter returns a new UserRouter instance
func NewRouter(ctrl *Controller, router *gin.Engine) *UserRouter {
	return &UserRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of user requests
func (r *UserRouter) RegisterRoutes() {
	r.router.GET("/user/:user_id/address", r.ctrl.handleGetUserAddress)
}
