package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/fomate-backend.git/controller/helper"
	"github.com/yebology/fomate-backend.git/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPurchasedContent(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("user_id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.GetError(c, "Error invalid user id format.")
	}

	contentIds, err := helper.GetPurchasedContentIds(c, ctx, objectId)
	if err != nil {
		return errors.GetError(c, "Error while retrieving content ids.")
	}

	contents, err := helper.GetContentByFilter(c, ctx, bson.M{"id": bson.M{"$in": contentIds}})
	if err != nil {
		return errors.GetError(c, "Error while retrieving contents.")
	}

	return c.Status(fiber.StatusOK).JSON(contents)

}

func GetUnpurchasedContent(c *fiber.Ctx) error {
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("user_id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.GetError(c, "Error invalid user id format.")
	}

	contentIds, err := helper.GetPurchasedContentIds(c, ctx, objectId)
	if err != nil {
		return errors.GetError(c, "Error while retrieving content ids.")
	}

	contents, err := helper.GetContentByFilter(c, ctx, bson.M{"id": bson.M{"$nin": contentIds}})
	if err != nil {
		return errors.GetError(c, "Error while retrieving contents.")
	}

	return c.Status(fiber.StatusOK).JSON(contents)

}
