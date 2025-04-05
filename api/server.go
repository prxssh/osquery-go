package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prxssh/osquery-go/api/osquery"
	"github.com/prxssh/osquery-go/internal/repo"
)

func StartServer(repo *repo.Repo) error {
	gin.SetMode(gin.ReleaseMode)
	corsOpts := cors.New(cors.Config{
		AllowHeaders:     []string{"*"},
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
	})

	router := gin.Default()
	router.Use(corsOpts)

	router.GET("/ping", heartbeat)
	initAPIServices(router, repo)

	return router.Run(":6969")
}

func initAPIServices(router *gin.Engine, repo *repo.Repo) {
	osqRoutes := router.Group("/v1")
	osqClient := osquery.NewOsqueryAPIService(repo)

	osqClient.InitRoutes(osqRoutes)
}

func heartbeat(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": "1.0.0",
		"data":    "pong",
	})
}
