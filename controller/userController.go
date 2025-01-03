package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/fomate-backend.git/controller/helper"
	"github.com/yebology/fomate-backend.git/database"
	"github.com/yebology/fomate-backend.git/errors"
	"github.com/yebology/fomate-backend.git/model"
	"github.com/yebology/fomate-backend.git/model/embedded"
	"github.com/yebology/fomate-backend.git/model/param"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNewUser(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user model.User
	err := c.BodyParser(&user)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	collection := database.GetDatabase().Collection("user")
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)

}

func PurchaseContent(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var purchasedContent embedded.PurchasedContent
	err := c.BodyParser(&purchasedContent)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	_, err = primitive.ObjectIDFromHex(purchasedContent.UserId.Hex())
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	_, err = primitive.ObjectIDFromHex(purchasedContent.ContentId.Hex())
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	collection := database.GetDatabase().Collection("purchased_content")
	res, err := collection.InsertOne(ctx, purchasedContent)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)

}

func AddSchedule(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var schedule embedded.Schedule
	err := c.BodyParser(&schedule)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	_, err = primitive.ObjectIDFromHex(schedule.UserId.Hex())
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	collection := database.GetDatabase().Collection("schedule")
	res, err := collection.InsertOne(ctx, schedule)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)

}

func PurchaseAllContent(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var lifetimeDeal param.LifetimeDeal
	err := c.BodyParser(&lifetimeDeal)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	objectId, err := primitive.ObjectIDFromHex(lifetimeDeal.UserId.Hex())
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	var contents []model.Content
	collection := database.GetDatabase().Collection("content")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return errors.GetError(c, err.Error())
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &contents)
	if err != nil {
		return errors.GetError(c, err.Error())
	}
	
	var purchasedContents []interface{}
	for _, con := range contents {
		purchasedContents = append(purchasedContents, embedded.PurchasedContent{
			UserId: objectId,
			ContentId: con.Id,
		})
	}

	collection = database.GetDatabase().Collection("purchased_content")
	res, err := collection.InsertMany(ctx, purchasedContents)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)

}

func UpdateUserHealth(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var updateHealth param.UpdateHealth
	err := c.BodyParser(&updateHealth)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	objectId, err := primitive.ObjectIDFromHex(updateHealth.UserId.Hex())
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	collection := database.GetDatabase().Collection("user")
	filter := bson.M{"_id":objectId}
	update := bson.M{"$set":bson.M{"health":updateHealth.NewHealth}}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
	
}

func GetLoginUser(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var login param.Login
	err := c.BodyParser(&login)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	users, err := helper.GetUsers(ctx, bson.M{})
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	var currentUser *model.User
	for _, u := range users {
		if u.Email == login.Email && u.Password == login.Password {
			currentUser = &u
			return c.Status(fiber.StatusOK).JSON(currentUser)
		}
	}

	return errors.GetError(c, "Invalid email or password.")

}

func GetUserSchedule(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("userId")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	var schedules []embedded.Schedule
	collection := database.GetDatabase().Collection("schedule")
	cursor, err := collection.Find(ctx, bson.M{"userId": objectId})
	if err != nil {
		return errors.GetError(c, err.Error())
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &schedules)
	if err != nil {
		return errors.GetError(c, err.Error())
	}
	if len(schedules) == 0 {
		return c.Status(fiber.StatusOK).JSON([]embedded.Schedule{})
	}

	return c.Status(fiber.StatusOK).JSON(schedules)

}

func GetUserHealth(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("userId")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	users, err := helper.GetUsers(ctx, bson.M{"_id":objectId})
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(users[0].Health)

}