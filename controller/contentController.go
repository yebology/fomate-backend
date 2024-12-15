package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/fomate-backend.git/controller/helper"
	"github.com/yebology/fomate-backend.git/errors"
	"github.com/yebology/fomate-backend.git/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPurchasedContent(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("userId")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	contentIds, err := helper.GetPurchasedContentIds(ctx, objectId)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	var contents []model.Content
	if len(contentIds) > 0 {
		contents, err = helper.GetContentByFilter(ctx, bson.M{"_id": bson.M{"$in": contentIds}})
		if err != nil {
			return errors.GetError(c, err.Error())
		}
	} else {
		contents = []model.Content{}
	}

	return c.Status(fiber.StatusOK).JSON(contents)


}

func GetUnpurchasedContent(c *fiber.Ctx) error {
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("userId")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	contentIds, err := helper.GetPurchasedContentIds(ctx, objectId)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	var contents []model.Content
	if len(contentIds) > 0 {
		contents, err = helper.GetContentByFilter(ctx, bson.M{"_id": bson.M{"$nin": contentIds}})
	} else {
		contents, err = helper.GetContentByFilter(ctx, bson.M{})
	}
	if err != nil {
		return errors.GetError(c, err.Error())
	}
	if len(contents) == 0 {
		return c.Status(fiber.StatusOK).JSON([]model.Content{})
	}

	return c.Status(fiber.StatusOK).JSON(contents)

}
