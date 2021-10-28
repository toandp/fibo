package apiv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"golang.org/x/crypto/bcrypt"
	"team1.asia/fibo/config"
	"team1.asia/fibo/db"
	"team1.asia/fibo/db/entity"
	"team1.asia/fibo/db/repository"
	"team1.asia/fibo/log"
)

type M struct{}

// Login is a function to authentication from database.
// @Summary The user authentication
// @Description The user authentication
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} M{}
// @Failure 400 {object} M{}
// @Failure 401 {object} M{}
// @Router /api/v1/login [post]
func Login(c *fiber.Ctx) error {
	user := c.Locals("user").(*entity.User)
	token := user.CreateJWTToken(config.App.JWT.Secret)

	c.Append("X-Access-Token", token.Hash)

	return c.JSON(fiber.Map{
		"data": user,
	})
}

// Validate the POST login request.
func ValidateLoginRequest(c *fiber.Ctx) error {
	var data entity.UserLoginForm

	if err := c.BodyParser(&data); err != nil {
		log.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	v := validate.Struct(data)

	if !v.Validate() {
		log.Error(v.Errors.String())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": v.Errors,
		})
	}

	repo := repository.New(db.ORM)
	user := repo.GetByUsername(data.Username)

	if user == nil {
		log.Error("User not found.")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not found.",
		})
	}

	match := ComparePasswordHash(data.Password, user.PasswordHash)

	if !match {
		log.Error("Invalid email or password.")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password.",
		})
	}

	c.Locals("user", user)

	return c.Next()
}

// Compares a bcrypt hashed password with user password.
func ComparePasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err != nil
}
