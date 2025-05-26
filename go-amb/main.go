package main

import (
	"amb/src/database"
	"amb/src/routes"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	// Initialize a new Fiber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return true // Разрешить все origins
		},
	}))

	routes.Setup(app)
	// Start the server on port 3000
	log.Fatal(app.Listen(":8000"))
}
