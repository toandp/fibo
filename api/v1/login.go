package apiv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"team1.asia/fibo/config"
	"team1.asia/fibo/db"
	"team1.asia/fibo/db/entity"
	"team1.asia/fibo/db/repository"
	"team1.asia/fibo/pkg"
)

func Login(c *fiber.Ctx) error {
	user := c.Locals("user").(*entity.User)

	token, _ := pkg.CreateJWTToken(user.Username, config.App.JWT.Secret)

	c.Append("X-Access-Token", token.Hash)

	return c.JSON(fiber.Map{
		"data": user,
	})
}

func ValidateLoginRequest(c *fiber.Ctx) error {
	var data entity.UserLoginForm

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	v := validate.Struct(data)

	if !v.Validate() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": v.Errors,
		})
	}

	repo := repository.New(db.ORM)

	user, err := repo.GetByUsername(data.Username)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	match, _ := pkg.CompareHash(data.Password, user.PasswordHash)

	if !match {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	c.Locals("user", user)

	return c.Next()
}
