package main

import (
	"github.com/nynrathod/uber-ride/api"
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	cfg "github.com/nynrathod/uber-ride/config"
	"github.com/nynrathod/uber-ride/pkg/services"
	ws "github.com/nynrathod/uber-ride/websocket"
)

func main() {
	// Initialize Fiber
	app := fiber.New()

	// Global middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // In production, restrict to trusted domains.
	}))
	app.Use(logger.New())

	// Initialize environment variables
	cfg.InitEnvConfigs()

	// Connect to PostgreSQL and obtain DB instance
	db := cfg.ConnectDB()

	// Setup WebSocket routes
	ws.SetupWebSocket(app, db)

	// Group API routes
	apiGroup := app.Group("/api")
	v1 := apiGroup.Group("/v1")

	// Use a wait group to ensure that service initialization completes before route installation.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		// Initialize application services (bundling repositories, usecases, etc.)
		serviceInitializer := services.NewAppServiceInitializer()
		appServices := serviceInitializer.InitializeAppServices()

		// Register API routes using the initialized services.
		api.InstallRouter(v1, appServices)
		wg.Done()
	}()
	wg.Wait()

	// Basic test route (optional)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start the Fiber server.
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
