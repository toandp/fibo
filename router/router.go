package router

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	apiv1 "team1.asia/fibo/api/v1"
	"team1.asia/fibo/config"
	"team1.asia/fibo/log"
)

// JWT authentication middleware.
// @param  cfg *config.AppConfig
// @return fiber.Handler
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

// Register the application routes.
// @param  app *fiber.App
// @return void
func RegisterRoutes(app *fiber.App) {
	group := app.Group("api")

	// V1
	registerV1Routes(group)

	// Handle not founds
	app.Use(func(c *fiber.Ctx) error {
		log.Zap.Info("404 not found.")
		return c.Status(404).JSON(fiber.Map{
			"error": "404 not found.",
		})
	})
}

// Register the routes V1.
// @param  group fiber.Router
// @return void
func registerV1Routes(group fiber.Router) {
	v1 := group.Group("v1")

	v1.Post("/login", apiv1.ValidateLoginRequest, apiv1.Login)
	v1.Post("/register", apiv1.ValidateRegisterRequest, apiv1.Register)
}
