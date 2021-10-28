package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	apiv1 "team1.asia/fibo/api/v1"
	"team1.asia/fibo/config"
	"team1.asia/fibo/log"
)

// JWT authentication.
func JWTAuthenticate(cfg *config.AppConfig) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(cfg.JWT.Secret),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(401).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	})
}

// Regsiter application routes.
func RegisterRoutes(app *fiber.App) {
	group := app.Group("api")

	// V1
	registerV1Routes(group)

	// Swagger router
	app.Get("/api-docs/*", swagger.Handler)

	// Handle not founds
	app.Use(func(c *fiber.Ctx) error {
		log.Error("404 not found.")
		return c.Status(404).JSON(fiber.Map{
			"error": "404 not found.",
		})
	})
}

// Register API V1 routes.
func registerV1Routes(group fiber.Router) {
	v1 := group.Group("v1")

	v1.Post("/login", apiv1.ValidateLoginRequest, apiv1.Login)
	v1.Post("/register", apiv1.ValidateRegisterRequest, apiv1.Register)
}
