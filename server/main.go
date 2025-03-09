package main

import (
	"github.com/nynrathod/uber-ride/pkg/services"
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/nynrathod/uber-ride/api"
	"github.com/nynrathod/uber-ride/config"
)

func main() {
	app := fiber.New()

	// Global middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(logger.New())

	// Initialize environment variables
	config.InitEnvConfigs()

	// Connect to PostgreSQL
	config.ConnectDB()

	// Group API routes
	apiGroup := app.Group("/api")
	v1 := apiGroup.Group("/v1")

	// Use a wait group to ensure service initialization completes before route installation
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		// Initialize application services (this bundles repositories, usecases, etc.)
		serviceInitializer := services.NewAppServiceInitializer()
		appServices := serviceInitializer.InitializeAppServices()

		// Register routes using the initialized services
		router.InstallRouter(v1, appServices)
		wg.Done()
	}()
	wg.Wait()

	// Basic test route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start the Fiber server
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
