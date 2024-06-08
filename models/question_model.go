package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	ID            primitive.ObjectID `json:"id,omitempty"`
	QuestionName  string             `json:"question"`
	Options       []string           `json:"options"`
	CorrectAnswer string             `json:"correct___answer"`
	Distractors   []string           `json:"distractors"`
}
