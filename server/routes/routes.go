package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/create", NewUrl)
	server.GET("/:id", GetUrl)
}
