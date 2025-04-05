package osquery

import "github.com/gin-gonic/gin"

func (osq *OsqueryAPIService) InitRoutes(router *gin.RouterGroup) {
	router.GET("/latest_data", osq.latestData)
}
