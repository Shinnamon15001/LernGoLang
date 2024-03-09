package main

import "github.com/gofiber/fiber/v2"

func main() {
	router := fiber.New()
	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(
			fiber.Map{
				"message": "Hello World!",
			},
		)
	})

	router.Listen(":8000")
}
