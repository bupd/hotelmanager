package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bupd/hotelmanager/db"
	"github.com/bupd/hotelmanager/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		panic(err)
	}

	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal("the db is dropping..")
	}

	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewMongoRoomStore(client, db.DBNAME, hotelStore)

	hotel := types.Hotel{
		Name:     "Kumarii",
		Location: "Singhii",
		Rooms:    []primitive.ObjectID{},
	}
	rooms := []types.Room{
		{
			Type:      types.SingleRoomType,
			BasePrice: 99.9,
		},
		{
			Type:      types.SingleRoomType,
			BasePrice: 99.9,
		},
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range rooms {
		r.HotelID = insertedHotel.ID
		insertedRoom, err := roomStore.InsertRoom(ctx, &r)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(insertedHotel)
		fmt.Println("--- above is hotel and below is room.")
		fmt.Println(insertedRoom, "this is room")
	}

	fmt.Println("seeding the database..")
}
