package osquery

import "github.com/gin-gonic/gin"

func (osq *OsqueryAPIService) InitRoutes(router *gin.RouterGroup) {
	router.GET("/latest-data", osq.latestData)
}
