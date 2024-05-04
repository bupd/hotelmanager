package db

import (
	"context"

	"github.com/bupd/hotelmanager/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DBNAME   = "hotelmanager"
	userColl = "users"
)

type UserStore interface {
	GetUserById(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	CreateUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context,filter bson.E, update bson.M) error
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(c *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: c,
		coll:   c.Database(DBNAME).Collection(userColl),
	}
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, filter bson.M, values bson.E) error {
  update:= bson.D{
    {
      "$set", bson.D{values},
    },
  } 
  _, err := s.coll.UpdateOne(ctx, filter, update)
	return nil
}

func (s *MongoUserStore) DeleteUser(ctx context.Context, userID string) error {
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

  // If you want check this later 
	_, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid})
  if err != nil {
    return err
  }
	return nil
}

func (s *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User
	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err // Changed to return nil instead of an empty slice
	}
	defer cur.Close(ctx) // Close the cursor once done

	// Iterate through the cursor and decode each document into a user
	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *MongoUserStore) GetUserById(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
