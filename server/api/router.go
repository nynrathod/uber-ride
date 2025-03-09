package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nynrathod/uber-ride/pkg/services"
)

// RequestBody is used for token generation.
type RequestBody struct {
	Email string `json:"email"`
}

// InstallRouter registers all User-related API routes.
func InstallRouter(app fiber.Router, services services.AppServices) {
	// Health-check or welcome route.
	app.Post("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello from User API",
		})
	})

	//// User Authentication & Profile routes.
	//app.Post("/auth/login", users.Login(services.UserService))
	//app.Post("/auth/register", middleware.Protected(), users.Register(services.UserService))
	//app.Post("/auth/verifyOtp", middleware.Protected(), users.VerifyOtp(services.UserService))
	//app.Post("/user/profile", middleware.Protected(), users.GetUser(services.UserService))
	//app.Post("/user/listusers", middleware.Protected(), users.ListUser(services.UserService))
	//
	//// Token generation route.
	//app.Post("/auth/gentoken", func(ctx *fiber.Ctx) error {
	//	var body RequestBody
	//	if err := ctx.BodyParser(&body); err != nil {
	//		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//			"status":  "error",
	//			"message": "Invalid request body",
	//		})
	//	}
	//	// Generate JWT token with email as a claim.
	//	additionalClaims := jwt.MapClaims{
	//		"email": body.Email,
	//	}
	//	token, err := utilities.GenerateJWT(additionalClaims)
	//	if err != nil {
	//		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//			"status":  "error",
	//			"message": "Failed to generate token",
	//		})
	//	}
	//	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
	//		"message": token,
	//	})
	//})
}
