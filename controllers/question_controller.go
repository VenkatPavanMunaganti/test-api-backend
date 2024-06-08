package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-mongo-api/configs"
	"github.com/gin-gonic/gin"
	"github.com/zeekhoks/test-api-backend/models"
	"github.com/zeekhoks/test-api-backend/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

var questionCollection *mongo.Collection = configs.GetCollection(configs.DB, "questions")

func GetDisplayQuestionsByTopicHandler() gin.HandlerFunc {

	return func(context *gin.Context) {

		params := context.Request.URL.Query()

		if params.Get("topic") == "" {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Topic not provided in URL",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"topic": params.Get("topic"),
		})
	}
}

func UploadQuestionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		file, _ := c.FormFile("questions_file")
		f, _ := file.Open()
		defer f.Close()
		content := make([]byte, file.Size)
		f.Read(content)
		err = json.Unmarshal(byteValue, &questionCollection)
		var question models.Question
		defer cancel()
		newQuestion := models.Question{
			ID:            primitive.NewObjectID(),
			QuestionName:  question.QuestionName,
			Options:       question.Options,
			CorrectAnswer: question.CorrectAnswer,
			Distractors:   question.Distractors,
		}
		result, err := questionCollection.InsertOne(ctx, newQuestion)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		log.Println(string(content))
		context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		c.JSON(http.StatusCreated)
	}
}
