package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bupd/hotelmanager/db"
	"github.com/bupd/hotelmanager/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		panic(err)
	}

	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewMongoRoomStore(client, db.DBNAME)

	hotel := types.Hotel{
		Name:     "Kumarii",
		Location: "Singhii",
	}
	room := types.Room{
		Type:      types.SingleRoomType,
		BasePrice: 99.9,
	}

	_ = room

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	room.HotelID = insertedHotel.ID
	insertedRoom, err := roomStore.InsertRoom(ctx, &room)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(insertedHotel)
	fmt.Println(insertedRoom)
	fmt.Println("seeding the database..")
}
