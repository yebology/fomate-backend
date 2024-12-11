package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/yebology/fomate-backend.git/database"
)

func main() {
	app := fiber.New()

	database.ConnectDatabase();
	defer database.DisconnectDatabase();

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
}