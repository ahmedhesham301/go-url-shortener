package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"urlshortener/db"
	"urlshortener/routes"
)

func main() {
	db.InitDB()
	defer db.Pool.Close()

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Use(cors.Default())

	server.Run(":8089")
}
