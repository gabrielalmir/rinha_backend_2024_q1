package config

import (
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/customer/entity"
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/customer/seeder"
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/db"
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/logger"
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/router"
)

var (
	log    *logger.Logger
	dbConn db.DBConn
)

func SetupApi() {
	// Initialize the logger
	log = logger.NewLogger("api")
	log.Info("Setting up the application ...")

	// Setup the database
	log.Info("Setting up the database ...")
	dbConfig := db.NewDBConfig("db", 5432, "postgres", "postgres", "postgres")
	dbConn = db.NewDBConn(dbConfig)
	err := dbConn.Connect()

	if err != nil {
		log.Errorf("Error opening the database connection: %s", err)
		return
	}

	// Migrate the database
	log.Info("Migrating the database ...")
	err = dbConn.GetDBConn().AutoMigrate(&entity.Customer{}, &entity.Transaction{})

	if err != nil {
		log.Errorf("Error migrating the database: %s", err)
		return
	}

	log.Info("Database setup complete")

	// Seed the database
	log.Info("Seeding the database ...")
	err = seeder.Seed(dbConn.GetDBConn())

	if err != nil {
		log.Errorf("Error seeding the database: %s", err)
		return
	}

	// End of database setup

	// Initialize routes
	log.Info("Setting up the routes ...")
	router.SetupRoutes(log)
	// End of routes setup

	log.Info("Application setup complete")
}
