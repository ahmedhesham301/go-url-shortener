package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/create", NewUrl)
	server.GET("/:id", GetUrl)
	server.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
