package main

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nynrathod/uber-ride/api"
	cfg "github.com/nynrathod/uber-ride/config"
	"github.com/nynrathod/uber-ride/pkg/services"
	"github.com/nynrathod/uber-ride/servertype"
	ws "github.com/nynrathod/uber-ride/websocket"
)

func main() {
	// Initialize Fiber
	app := fiber.New()

	// Global middleware
	app.Use(cors.New(cors.Config{AllowOrigins: "*"})) // In production, restrict to trusted domains.
	app.Use(logger.New())

	// Initialize environment variables
	env := cfg.InitEnvConfigs()

	// Connect to PostgreSQL and obtain DB instance
	db := cfg.ConnectDB()

	// Setup WebSocket routes
	ws.SetupWebSocket(app, db)

	// Group API routes
	apiGroup := app.Group("/api")
	v1 := apiGroup.Group("/v1")

	// Use a wait group to ensure service initialization before route installation.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		serviceInitializer := services.NewAppServiceInitializer()
		appServices := serviceInitializer.InitializeAppServices()
		api.InstallRouter(v1, appServices)
		wg.Done()
	}()
	wg.Wait()

	// Check environment variables
	httpsEnabled := env.HTTPS == "true"

	// Run in PRODUCTION with Let's Encrypt HTTPS
	if env.ENV == "PROD" {
		servertype.StartProductionServer(app)
	} else if httpsEnabled {
		servertype.StartLocalHttpsServer(app)
	} else {
		servertype.StartLocalHttpServer(app)
	}
}
