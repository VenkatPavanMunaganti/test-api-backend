package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/zeekhoks/test-api-backend/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	"os"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var uri = goDotEnvVariable("MONGODB_URI")

var mongoClient *mongo.Client

func init() {
	if err := connectToMongodb(); err != nil {
		log.Fatal("Could not connect to MongoDB")
	} else {
		log.Println("Connected to MongoDB")
	}
}

func main() {

	router := routes.GetRouter()

	err := router.Run(":8000")

	if err != nil {
		fmt.Printf("Fatal error has occured: %v\n", err)
	}
}

func connectToMongodb() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	mongoClient = client
	return err
}
