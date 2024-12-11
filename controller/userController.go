package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/fomate-backend.git/database"
	"github.com/yebology/fomate-backend.git/errors"
	"github.com/yebology/fomate-backend.git/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetUser(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("user_id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.GetError(c, "Error invalid id format")
	}

	var user model.User
	collection := database.GetDatabase().Collection("user")
	err = collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return errors.GetError(c, "Error while decoding user")
	}

	return c.Status(fiber.StatusOK).JSON(user)

}