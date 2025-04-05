package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() error {
	gin.SetMode(gin.ReleaseMode)
	corsOpts := cors.New(cors.Config{
		AllowHeaders:     []string{"*"},
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
	})

	router := gin.Default()
	router.Use(corsOpts)

	router.GET("/ping", heartbeat)
	initAPIServices(router)

	return router.Run(":6969")
}

func initAPIServices(router *gin.Engine) {
}

func heartbeat(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": "1.0.0",
		"data":    "pong",
	})
}
