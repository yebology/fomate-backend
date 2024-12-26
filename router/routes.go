package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/fomate-backend.git/controller"
)

func SetUp(app *fiber.App) {

	// done check postman
	app.Get("/api/get_all_app", controller.GetAllApp)

	// done check postman
	app.Get("/api/get_health/:userId", controller.GetUserHealth)

	// done check postman
	app.Get("/api/get_schedule/:userId", controller.GetUserSchedule)

	// done check postman
	app.Get("/api/get_purchased_content/:userId", controller.GetPurchasedContent)

	// done check postman
	app.Get("/api/get_unpurchased_content/:userId", controller.GetUnpurchasedContent)

	// done check postman
	app.Post("/api/get_user", controller.GetLoginUser)

	// done check postman
	app.Post("/api/create_user", controller.CreateNewUser)

	// done check postman
	app.Post("/api/purchase_content", controller.PurchaseContent)

	// done check postman
	app.Post("/api/purchase_all_content", controller.PurchaseAllContent)

	// done check postman
	app.Post("/api/add_schedule", controller.AddSchedule)

	// done check postman
	app.Post("/api/update_health", controller.UpdateUserHealth)

}