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

func GetAllApp(c *fiber.Ctx) error {
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var apps []model.App

	collection := database.GetDatabase().Collection("app")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return errors.GetError(c, err.Error())
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &apps)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(apps)

}