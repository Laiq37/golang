package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	//loading env file
	err := godotenv.Load(".env")

	//error handling if failed to load
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//getting mongodb_url
	MongoDb := os.Getenv("MONGODB_URL")

	//setting up mongodb client to use database
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))

	//error handling
	if err != nil {
		log.Fatal(err)
	}

	//getting context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//releasing resources
	defer cancel()

	//connecting to db
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	return client
}

// creating DBinstance
var Client *mongo.Client = DBinstance()

// func for opening and getting collection
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("jwt_auth").Collection(collectionName)
	return collection
}
