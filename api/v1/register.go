package apiv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"golang.org/x/crypto/bcrypt"
	"team1.asia/fibo/config"
	"team1.asia/fibo/db"
	"team1.asia/fibo/db/entity"
	"team1.asia/fibo/db/repository"
)

// POST /api/v1/register handler.
// @param  c *fiber.Ctx
// @return error
func Register(c *fiber.Ctx) error {
	user := c.Locals("user").(*entity.User)
	repo := c.Locals("repository").(repository.UserRepositoryInterface)

	repo.Create(user)

	token := user.CreateJWTToken(config.App.JWT.Secret)

	c.Append("X-Access-Token", token.Hash)

	return c.JSON(fiber.Map{
		"data": user,
	})
}

// Validate the POST register request.
// @param  c *fiber.Ctx
// @return error
func ValidateRegisterRequest(c *fiber.Ctx) error {
	var data entity.UserRegisterForm

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

	record := repo.GetByUsername(data.Username)

	if record != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "User already exists",
		})
	}

	var user entity.User

	user.Username = data.Username
	user.PasswordHash = CreatePasswordHash(data.Password)

	c.Locals("user", &user)
	c.Locals("repository", repo)

	return c.Next()
}

// Create the password hash.
// @param  password string
// @return string
func CreatePasswordHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		panic(err)
	}

	return string(hash)
}
