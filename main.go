package main

import (
	"context"
	"flag"

	"github.com/bupd/hotelmanager/api"
	"github.com/bupd/hotelmanager/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dburi    string = "mongodb://localhost:27017"
	dbname   string = "hotelmanager"
	userColl string = "users"
)

var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		panic(err)
	}

	listenAddr := flag.String("listenAddr", ":4444", "The listenAddr of the api Server")
	flag.Parse()

	app := fiber.New(config)

	apiv1 := app.Group("/api/v1")
	apiv1.Get("/healthz", handleHealthz)

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Get("/users", userHandler.HandleGetUsers)

	app.Get("/", handleMain)
	app.Listen(*listenAddr)
}

func handleMain(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msgs": "hello I am running in the server"})
}

func handleHealthz(c *fiber.Ctx) error {
	return c.JSON(map[string]int{"status": 200})
}
