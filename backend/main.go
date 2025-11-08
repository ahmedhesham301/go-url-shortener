package main

import (
	"urlshortener/db"
	"urlshortener/middleware"
	"urlshortener/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	defer db.Pool.Close()

	server := gin.Default()
	server.Use(cors.Default())
	server.Use(middleware.RequestLatency)
	routes.RegisterRoutes(server)

	server.Run(":8089")
}
