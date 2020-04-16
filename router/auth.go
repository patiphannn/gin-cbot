package router

import (
	"github.com/gin-gonic/gin"
	"github.com/polnoy/gin-cbot/controller"
)

// Auth defined user router
func Auth(r *gin.Engine) *gin.Engine {
	controller := new(controller.Auth)

	g := r.Group("/auth")
	g.POST("/signup", controller.Signup)
	g.POST("/signin", controller.Signin)
	g.POST("/refresh/token", controller.Refresh)

	return r
}
