package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func Connection(URI string) *mongo.Client {

	ctx, cancelCtx := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancelCtx()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		log.Println(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return client
}
