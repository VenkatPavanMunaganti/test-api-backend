package main

import (
	"fmt"
	"github.com/zeekhoks/test-api-backend/routes"
)

func main() {

	router := routes.GetRouter()

	err := router.Run(":8000")

	if err != nil {
		fmt.Printf("Fatal error has occured: %v\n", err)
	}
}
