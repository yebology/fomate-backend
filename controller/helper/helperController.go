package helper

import (
	"context"

	"github.com/yebology/fomate-backend.git/database"
	"github.com/yebology/fomate-backend.git/model"
	"github.com/yebology/fomate-backend.git/model/embedded"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPurchasedContentIds(ctx context.Context, userId primitive.ObjectID) ([] primitive.ObjectID, error) {

	var purchasedContents []embedded.PurchasedContent
	collection := database.GetDatabase().Collection("purchased_content")
	cursor, err := collection.Find(ctx, bson.M{"userId": userId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &purchasedContents)
	if err != nil {
		return nil, err
	}

	var contentIds []primitive.ObjectID
	for _, pc := range purchasedContents {
		contentIds = append(contentIds, pc.ContentId)
	}

	return contentIds, nil

}

func GetContentByFilter(ctx context.Context, filter bson.M) ([]model.Content, error) {

	var contents []model.Content
	collection := database.GetDatabase().Collection("content")
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &contents)
	if err != nil {
		return nil, err
	}

	return contents, nil

}

func GetUsers(ctx context.Context, filter bson.M) ([]model.User, error) {

	var users []model.User
	collection := database.GetDatabase().Collection("user")
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err 
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil

}
