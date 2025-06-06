package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/url", NewUrl)
	server.GET("/:hash", GetUrl)
}
