package middlewares

import (
	"hims/pkg/config"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var cfg, _ = config.LoadConfigs()

func Auth(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cfg.SECRET_KEY)},
		ContextKey: "jwt",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "true",
				"message": err.Error(),
			})
		},
	})(ctx)
}
