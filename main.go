package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mongo-golang/controllers"
)

func main() {
	// Declare and initialize the client
	client := getSession()
	uc := controllers.NewUserController(client)

	// Set up router
	r := httprouter.New()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.Create)
	r.DELETE("/user/:id", uc.DeleteUser)

	fmt.Println("Server running at :9000")
	http.ListenAndServe(":9000", r)
}

func getSession() *mongo.Client {
	// MongoDB connection string
	uri := "mongodb+srv://sauravkum420:3XmFmOS2yLCVCLIP@cluster0.i1zp1.mongodb.net/"

	// Create a new client
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("Failed to create MongoDB client:", err)
		panic(err)
	}

	// Set up a context with a timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Connect to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		panic(err)
	}

	// Test the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("MongoDB ping failed:", err)
		panic(err)
	}

	fmt.Println("Connected to MongoDB successfully!")
	return client
}
// 3XmFmOS2yLCVCLIP
