package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/fomate-backend.git/controller"
)

func SetUp(app *fiber.App) {

	app.Get("/api/get_user/:user_id", controller.CreateNewUser)

	app.Get("/api/get_all_content", )

	app.Post("/api/create_new_user", )

	app.Post("/api/buy_content/:user_id/:content_id", )

	app.Post("/api/buy_all_content/:user_id", )
}