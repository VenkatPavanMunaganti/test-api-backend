package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/zeekhoks/test-api-backend/routes"
	"github.com/zeekhoks/test-api-backend/services"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("No .env file available")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatalln("Mongodb URI string not found")
	}

	// connecting to the database
	err = services.ConnectToMongo(uri)
	if err != nil {
		log.Fatalln("Failed to connect to MongoDB")
	} else {
		log.Println("Connected to DB")
	}
}

func main() {
	DBCon := services.GetConnection()
	router := routes.GetRouter(DBCon)

	err := router.Run(":8000")

	if err != nil {
		fmt.Printf("Fatal error has occured: %v\n", err)
	}
}
