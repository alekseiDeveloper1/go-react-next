package main

import (
    "log"

    "github.com/gofiber/fiber/v3"
)

func main() {
    _, err := gorm.Open(mysql.Open("root:root@tcp(db:3306)/amb"), &gorm.Open())

    if err != nil {
        panic("Could not connect with the database!")
    }
    // Initialize a new Fiber app
    app := fiber.New()

    // Define a route for the GET method on the root path '/'
    app.Get("/", func(c fiber.Ctx) error {
        // Send a string response to the client
        return c.SendString("Hello, World 👋!")
    })

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}