package errors

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetError(c *fiber.Ctx, err string) error {
	fmt.Println("Error : ", err)
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
}