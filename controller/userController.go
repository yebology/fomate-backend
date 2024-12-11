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
		return errors.GetError(c, "Error while insert new data.")
	}

	return c.Status(fiber.StatusOK).JSON(res)

}

func PurchaseContent(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var purchasedContent model.PurchasedContent
	err := c.BodyParser(&purchasedContent)
	if err != nil {
		return errors.GetError(c, "Error while parsing data.")
	}

	_, err = primitive.ObjectIDFromHex(purchasedContent.UserId.Hex())
	if err != nil {
		return errors.GetError(c, "Error invalid user id format.")
	}

	_, err = primitive.ObjectIDFromHex(purchasedContent.ContentId.Hex())
	if err != nil {
		return errors.GetError(c, "Error invalid content id format.")
	}

	collection := database.GetDatabase().Collection("purchased_content")
	res, err := collection.InsertOne(ctx, purchasedContent)
	if err != nil {
		return errors.GetError(c, "Error while insert new data.")
	}

	return c.Status(fiber.StatusOK).JSON(res)

}

func AddSchedule(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var schedule model.Schedule
	err := c.BodyParser(&schedule)
	if err != nil {
		return errors.GetError(c, "Error while parsing data.")
	}

	_, err = primitive.ObjectIDFromHex(schedule.AppId.Hex())
	if err != nil {
		return errors.GetError(c, "Error invalid app id format.")
	}

	_, err = primitive.ObjectIDFromHex(schedule.UserId.Hex())
	if err != nil {
		return errors.GetError(c, "Error invalid user id format.")
	}

	collection := database.GetDatabase().Collection("schedule")
	res, err := collection.InsertOne(ctx, schedule)
	if err != nil {
		return errors.GetError(c, "Error while insert new data.")
	}

	return c.Status(fiber.StatusOK).JSON(res)

}


func GetUser(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("user_id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.GetError(c, "Error invalid user id format.")
	}

	var user model.User
	collection := database.GetDatabase().Collection("user")
	err = collection.FindOne(ctx, bson.M{"id": objectId}).Decode(&user)
	if err != nil {
		return errors.GetError(c, "Error while decoding user.")
	}

	return c.Status(fiber.StatusOK).JSON(user)

}