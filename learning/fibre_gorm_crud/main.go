package main

import (
	"fiber-gorm-crud-api/database"
	"fiber-gorm-crud-api/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	LoadAppConfig()

	//initailizing Fiber router/App instance
	app := fiber.New()

	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	//defining routes
	app.Get("/dogs", handlers.GetDogs)
	app.Get("/dogs/:id", handlers.GetDog)
	app.Post("/dogs", handlers.AddDog)
	app.Put("/dogs/:id", handlers.UpdateDog)
	app.Delete("/dogs/:id", handlers.RemoveDog)

	log.Fatal(app.Listen(":3000"))
}
