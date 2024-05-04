package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/bupd/hotelmanager/api"
	"github.com/bupd/hotelmanager/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dburi    string = "mongodb://localhost:27017"
	dbname   string = "hotelmanager"
	userColl string = "users"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		panic(err)
	}
	coll := client.Database(dbname).Collection(userColl)

	user := types.User{
		FirstName: "Prasanth",
		LastName:  "bupd",
	}
	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	var kumar types.User
	if err := coll.FindOne(ctx, bson.M{}).Decode(&kumar); err != nil {
		log.Fatal(err)
	}

	fmt.Println("the user returned.", kumar)

	fmt.Println("Showing the client and collection", client, coll)

	listenAddr := flag.String("listenAddr", ":4444", "The listenAddr of the api Server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")
	apiv1.Get("/healthz", handleHealthz)
  apiv1.Get("/user/:id", api.HandleGetUser)
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
