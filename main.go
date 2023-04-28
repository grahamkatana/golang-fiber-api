package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/grahamkatana/fiber-api/database"
	"github.com/grahamkatana/fiber-api/routes"
)

func setupRoutes(app *fiber.App) {
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Patch("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id",routes.DeleteUser)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
