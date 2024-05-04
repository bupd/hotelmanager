package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bupd/hotelmanager/db"
	"github.com/bupd/hotelmanager/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testdburi    string = "mongodb://localhost:27017"
	testdbname   string = "hotelmanager-test"
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
		UserStore: db.NewMongoUserStore(client, testdbname),
	}
}

func TestPotta(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandlePostUser)

	// Define multiple user creation parameters
	params := types.CreateUserParams{
		Email: "test@test.com", FirstName: "Test", LastName: "User", Password: "password2",
	}
	// Add more users as needed

	// Convert user creation parameters to JSON
	b, _ := json.Marshal(params)

	// Create HTTP request with JSON payload
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	// Send request to the application
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	// Check response status
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status %d; got %d", http.StatusOK, resp.StatusCode)
	}

	// Decode response body to get user details
	var user types.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if len(user.ID) == 0 {
		t.Errorf("Expected userID to be set")
	}
	if len(user.EncryptedPassword) > 0 {
		t.Errorf("Fatal Error returning Password!!!")
	}
	if user.FirstName != params.FirstName {
		t.Errorf("Expected %s but got %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("Expected %s but got %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("Expected %s but got %s", params.Email, user.Email)
	}

	defer resp.Body.Close()
}
