package router

import (
	"github.com/gofiber/fiber/v2"
	apiv1 "team1.asia/fibo/api/v1"
)

func RegisterRoutes(app *fiber.App) {
	group := app.Group("api")

	// V1
	registerV1Routes(group)
}

func registerV1Routes(group fiber.Router) {
	v1 := group.Group("v1")

	v1.Post("/login", apiv1.ValidateLoginRequest, apiv1.Login)
	v1.Post("/register", apiv1.ValidateRegisterRequest, apiv1.Register)
}
