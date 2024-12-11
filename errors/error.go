package errors

import "github.com/gofiber/fiber/v2"

func GetError(c *fiber.Ctx, err string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
}