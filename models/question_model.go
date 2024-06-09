package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	QuestionName  string             `json:"question" bson:"question"`
	Options       []string           `json:"options" bson:"options"`
	CorrectAnswer string             `json:"-" bson:"correct_answer"`
	Distractors   []string           `json:"-" bson:"distractors"`
}
