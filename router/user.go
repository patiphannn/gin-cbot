package router

import (
	"github.com/gin-gonic/gin"
	"github.com/polnoy/gin-cbot/controller"
)

// User defined user router
func User(r *gin.Engine) *gin.Engine {
	controller := new(controller.User)

	g := r.Group("/user")
	g.GET("", controller.Gets)
	g.GET("/:_id", controller.Get)
	g.POST("/", controller.Create)
	g.PUT("/:_id", controller.Update)
	g.DELETE("/:_id", controller.DeleteByID)

	return r
}
