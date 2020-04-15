package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/polnoy/gin-cbot/common"
	"github.com/polnoy/gin-cbot/router"
)

func init() {
	log.Println("main init: ")
	common.ConnectDb()
}

func main() {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// User router
	r = router.User(r)

	r.Run(":8080")
}
