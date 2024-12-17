package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Hello Fiber Completed")
		return c.SendString("Hello Fiber!")
	})

	app.Listen(":3000")
}
