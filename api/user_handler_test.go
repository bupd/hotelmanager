package api

import (
	"context"
	"testing"

	"github.com/bupd/hotelmanager/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testdburi    string = "mongodb://localhost:27017"
	testdbname   string = "hotelmanager"
	testuserColl string = "users"
)

type testdb struct {
	db.UserStore
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))
	if err != nil {
		panic(err)
	}
	return &testdb{
		UserStore: db.NewMongoUserStore(client),
	}
}

func TestPotta(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)
	t.Fail()
}
