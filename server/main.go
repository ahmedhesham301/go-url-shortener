package main

import (
	"urlshortener/db"
	"urlshortener/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	defer db.Pool.Close()

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Use(cors.Default())

	server.Run(":8089")
}
