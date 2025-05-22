package main

import (
	"amb/src/database"
	"amb/src/routes"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	// Initialize a new Fiber app
	app := fiber.New()

	routes.Setup(app)
	// Start the server on port 3000
	log.Fatal(app.Listen(":8000"))
}
