package main

import (
	"myapp/config"
	"myapp/routes"
	"myapp/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Log startup info
	utils.Info("Starting the server...")

	// Create router
	router := gin.Default()

	// Initialize routes
	routes.InitRoutes(router)

	// Start server
	port := config.GetEnv("PORT", "8080")
	utils.Info("Server running on port " + port)
	router.Run(":" + port)
}
