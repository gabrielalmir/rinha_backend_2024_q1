package main

import (
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/config"
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/logger"
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/router"
)

func main() {
	// Initialize the logger
	log := logger.GetLogger("API")

	// Initialize configurations
	config.Init(log)

	// Initialize routes
	log.Info("Setting up the routes ...")
	router.SetupRoutes(log)
}
