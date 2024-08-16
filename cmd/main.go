package main

import (
	"log"

	"context"
	"crud-hexagonal/config"
	handlers "crud-hexagonal/internal/adapter/handlers/http"
	"crud-hexagonal/internal/adapter/logs"
	"crud-hexagonal/internal/adapter/repository"
	"crud-hexagonal/internal/core/service"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create a context with a timeout for MongoDB operations
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Initialize MongoDB connection
	db, err := repository.NewMongoDB(ctx, cfg.MongoDB.DSN, cfg.MongoDB.DBName)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Create product repository
	productRepo := repository.NewProductsRepository(db.Database)

	// Create service
	service := service.NewProductsService(productRepo)

	// Create Fiber app
	app := fiber.New()

	// Register routes
	handlers.NewRouter(app, service)

	// Start server
	logs.NewLogger(app)
	log.Fatal(app.Listen(":" + cfg.App.Port))
}
