package middlewares

import (
	"hims/pkg/config"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
)

var cfg *config.Config

// SetConfig sets the middleware config at runtime (avoid package init loads)
func SetConfig(c *config.Config) {
	cfg = c
}

// verifies the token if there is any
// will extract the claims and set it in context key
func Auth(ctx fiber.Ctx) error {
	if cfg == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "true",
			"message": "server configuration not loaded",
		})
	}
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cfg.SECRET_KEY)},
		ErrorHandler: func(ctx fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "true",
				"message": err.Error(),
			})
		},
	})(ctx)
}
