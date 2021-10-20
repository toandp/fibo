package pkg

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"team1.asia/fibo/config"
)

type AccessToken struct {
	Hash   string `json:"access_token"`
	Expire int64  `json:"expires_in"`
}

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

func CreateJWTToken(user string, secret string, expires ...int64) (*AccessToken, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	expire := config.App.JWT.Expire

	if len(expires) > 0 {
		expire = expires[0]
	}

	expiresIn := time.Now().Add(time.Duration(expire) * time.Second).Unix()

	claims["username"] = user
	claims["exp"] = expiresIn

	tokenHash, err := token.SignedString([]byte(secret))

	if err != nil {
		return nil, err
	}

	return &AccessToken{
		Hash:   tokenHash,
		Expire: expiresIn,
	}, nil
}
