package main

import "github.com/gofiber/fiber/v2"
import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	listenAddr := flag.String("listenAddr", ":4444", "The listenAddr of the api Server")
	flag.Parse()
	app.Get("/", handleMain)
	app.Listen(*listenAddr)
}

func handleMain(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msgs": "hello I am running in the server"})
}
