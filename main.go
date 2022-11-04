package main

import (
	//fiber
	"meditation/controllers"
	"meditation/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	database.DatabaseMain()

	app.Static("/storage/music", "./assets/music")
	app.Get("/music", controllers.GetMusic)

	app.Listen(":3000")

}
