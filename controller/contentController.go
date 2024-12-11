package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/fomate-backend.git/database"
	"github.com/yebology/fomate-backend.git/errors"
	"github.com/yebology/fomate-backend.git/model"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllContent(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var contents []model.Content
	collection := database.GetDatabase().Collection("content")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return errors.GetError(c, "Error while find content id.")
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, contents)
	if err != nil {
		return errors.GetError(c, "Error while decoding data.")
	}

	return c.Status(fiber.StatusOK).JSON(contents)
}