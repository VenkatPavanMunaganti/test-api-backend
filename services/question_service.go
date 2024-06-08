package services

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var questionCollection *mongo.Collection
var COLLECTION = "questions"
