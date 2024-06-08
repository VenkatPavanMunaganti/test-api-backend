package services

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var client *mongo.Client
var questionCollection *mongo.Collection
var COLLECTION = "questions"

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	if questionCollection != nil {
		return questionCollection
	}
	questionCollection := client.Database("Pearson").Collection(collectionName)
	return questionCollection
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func GetClient() *mongo.Client {
	uri := goDotEnvVariable("MONGODB_URI")
	//getting context
	if client != nil {
		return client
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//getting client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln(err)
	}
	return client
}
