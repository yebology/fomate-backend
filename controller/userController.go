package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/fomate-backend.git/database"
	"github.com/yebology/fomate-backend.git/errors"
	"github.com/yebology/fomate-backend.git/model"
)


func CreateNewUser(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user model.User
	err := c.BodyParser(&user)
	if err != nil {
		return errors.GetError(c, "Error while parsing data.")
	}

	collection := database.GetDatabase().Collection("user")
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return errors.GetError(c, "Error while insert new user data.")
	}

	return c.Status(fiber.StatusOK).JSON(res)

}

// func GetUser(c *fiber.Ctx) error {

// }