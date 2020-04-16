package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/polnoy/gin-cbot/router"
)

func main() {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Auth router
	r = router.Auth(r)

	// User router
	r = router.User(r)

	r.Run(":8080")
}
