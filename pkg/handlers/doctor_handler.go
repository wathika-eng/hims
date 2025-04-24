package handlers

import "github.com/gofiber/fiber/v2"

func TestAPI(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"data": "Hello, World!",
	})
}
