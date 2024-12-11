package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/fomate-backend.git/controller"
)

func SetUp(app *fiber.App) {

	app.Get("/api/get_user/:user_id", controller.GetUser)

	app.Get("/api/get_all_app", controller.GetAllApp)

	app.Get("/api/get_schedule/:user_id", controller.GetUserSchedule)

	app.Get("/api/get_purchased_content/:user_id", controller.GetPurchasedContent)

	app.Get("/api/get_unpurchased_content/:user_id", controller.GetUnpurchasedContent)

	app.Post("/api/create_user", controller.CreateNewUser)

	app.Post("/api/purchase_content", controller.PurchaseContent)

	app.Post("/api/purchase_all_content/:user_id", controller.PurchaseAllContent)

	app.Post("/api/add_schedule", controller.AddSchedule)

}