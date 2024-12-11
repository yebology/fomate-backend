package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/fomate-backend.git/controller"
)

func SetUp(app *fiber.App) {

	app.Get("/api/get_user/:user_id", controller.GetUser)

	app.Get("/api/get_all_content", controller.GetAllContent)

	app.Post("/api/create_user", controller.CreateNewUser)

	app.Post("/api/purchase_content/:user_id/:content_id", )

	app.Post("/api/purchase_all_content/:user_id", )

	app.Post("/api/add_schedule/:user_id/:app_id", )
	
}