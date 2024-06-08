package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zeekhoks/test-api-backend/models"
	"github.com/zeekhoks/test-api-backend/services"
	"log"
	"net/http"
)

//var questionCollection = services.GetCollection(DB, "questions")

func GetDisplayQuestionsByTopicHandler() gin.HandlerFunc {

	return func(context *gin.Context) {
		DB := services.GetConnection()
		questionsCollection := services.GetCollection(DB, "questions")

		if questionsCollection != nil {
			log.Println("collection found")
		}

		params := context.Request.URL.Query()

		if params.Get("topic") == "" {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Topic not provided in URL",
			})
			return
		}
		err := DB.Ping(context, nil)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"topic":      params.Get("topic"),
				"connection": false,
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"topic":      params.Get("topic"),
				"connection": true,
			})
		}
	}
}

func UploadQuestionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		DB := services.GetConnection()
		questionsCollection := services.GetCollection(DB, "questions")

		file, _ := c.FormFile("questions_file")
		f, _ := file.Open()
		defer f.Close()
		content := make([]byte, file.Size)
		_, err := f.Read(content)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Server error. Please try again later",
			})
			return
		}
		var questions []models.Question
		err = json.Unmarshal(content, &questions)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Server error. Please try again later",
			})
			return
		}

		for i := 0; i < len(questions); i++ {
			fmt.Printf("%v\n", questions[i])
		}
		var interfaces []interface{}
		for _, question := range questions {
			interfaces = append(interfaces, question)
		}

		res, err := questionsCollection.InsertMany(c, interfaces)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Server error. Please try again later",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"insertedIDs": fmt.Sprintf("%v", res.InsertedIDs),
		})

	}
}
