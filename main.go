package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/", handleMain)
	app.Listen(":4444")
}

func handleMain(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msgs": "hello I am running in the server"})
}
