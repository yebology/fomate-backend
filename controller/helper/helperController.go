package helper

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/fomate-backend.git/database"
	"github.com/yebology/fomate-backend.git/errors"
	"github.com/yebology/fomate-backend.git/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPurchasedContentIds(c *fiber.Ctx, ctx context.Context, userId primitive.ObjectID) ([] primitive.ObjectID, error) {

	var purchasedContents []model.PurchasedContent
	collection := database.GetDatabase().Collection("purchased_content")
	cursor, err := collection.Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		return nil, errors.GetError(c, "Error while find purchased content id.")
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &purchasedContents)
	if err != nil {
		return nil, errors.GetError(c, "Error while decoding data.")
	}

	var contentIds []primitive.ObjectID
	for _, pc := range purchasedContents {
		contentIds = append(contentIds, pc.ContentId)
	}

	return contentIds, nil

}

func GetContentByFilter(c *fiber.Ctx, ctx context.Context, filter bson.M) ([]model.Content, error) {
	
	var contents []model.Content
	collection := database.GetDatabase().Collection("content")
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.GetError(c, "Error while find purchased content.")
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &contents)
	if err != nil {
		return nil, errors.GetError(c, "Error while decoding data.")
	}

	return contents, nil
}

