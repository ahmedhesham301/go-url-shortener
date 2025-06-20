package main

import (
	"urlshortener/db"
	"urlshortener/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8089")
}
