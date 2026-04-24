package main

import (
	"log"
	"os"

	"lemonade/config"
	"lemonade/middleware"
	"lemonade/routes"
	"lemonade/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize logger
	logger := utils.NewLogger()
	logger.Info("Starting Lemonade API Server")

	// Set Gin mode
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "debug"
	}
	gin.SetMode(ginMode)

	// Load configuration
	cfg := config.LoadConfig()
	logger.Info("Configuration loaded", map[string]interface{}{
		"port": cfg.Port,
		"env":  cfg.Env,
	})

	// Create Gin router
	router := gin.Default()

	// Apply middleware
	router.Use(middleware.CORSMiddleware())

	// Setup routes
	routes.SetupRoutes(router)

	// Start server
	addr := ":" + cfg.Port
	logger.Info("Starting server on " + addr)

	if err := router.Run(addr); err != nil {
		logger.Error("Failed to start server", map[string]interface{}{
			"error": err.Error(),
		})
		os.Exit(1)
	}
}
