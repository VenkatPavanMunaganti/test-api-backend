package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/zeekhoks/test-api-backend/routes"
	"github.com/zeekhoks/test-api-backend/services"
	"golang.org/x/net/context"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("No .env file available")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatalln("Mongodb URI string not found")
	}
	// connecting to the database
	client := services.Connection(uri)
	if client == nil {
		log.Fatalln("Failed to connect to MongoDB")
	} else {
		log.Println("Connected to MongoDB")
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatalln(err)
			return
		}
	}()

	router := routes.GetRouter()

	err = router.Run(":8000")

	if err != nil {
		fmt.Printf("Fatal error has occured: %v\n", err)
	}
}
