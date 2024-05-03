package main

import (
	"flag"

	"github.com/bupd/hotelmanager/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	listenAddr := flag.String("listenAddr", ":4444", "The listenAddr of the api Server")
	flag.Parse()

	apiv1 := app.Group("/api/v1")
	apiv1.Get("/healthz", handleHealthz)
	apiv1.Get("/user", api.HandleGetUser)
	apiv1.Get("/users", api.HandleGetUsers)

	app.Get("/", handleMain)
	app.Listen(*listenAddr)
}

func handleMain(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msgs": "hello I am running in the server"})
}

func handleHealthz(c *fiber.Ctx) error {
	return c.JSON(map[string]int{"status": 200})
}
