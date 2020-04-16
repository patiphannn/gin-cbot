package router

import (
	"github.com/gin-gonic/gin"
	"github.com/polnoy/gin-cbot/common"
	"github.com/polnoy/gin-cbot/controller"
)

// User defined user router
func User(r *gin.Engine) *gin.Engine {
	controller := new(controller.User)

	g := r.Group("/user")
	g.GET("", common.TokenAuthMiddleware(), controller.Gets)
	g.GET("/:_id", common.TokenAuthMiddleware(), controller.Get)
	g.POST("/", common.TokenAuthMiddleware(), controller.Create)
	g.PUT("/:_id", common.TokenAuthMiddleware(), controller.Update)
	g.DELETE("/:_id", common.TokenAuthMiddleware(), controller.DeleteByID)

	return r
}
