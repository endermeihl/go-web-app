package main

import (
	"go-web-app/config"
	"go-web-app/routes"
	"go-web-app/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configurations
	config.LoadConfig()

	// Initialize MySQL and Redis connections
	utils.ConnectDatabase()
	utils.ConnectRedis()

	// Set up the router
	r := gin.Default()
	routes.RegisterRoutes(r)

	// Start the server
	r.Run(":8080")
}
