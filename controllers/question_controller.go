package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zeekhoks/test-api-backend/models"
	"github.com/zeekhoks/test-api-backend/services"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func GetDisplayQuestionsByTopicHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		//get params from the request
		params := context.Request.URL.Query()
		if params.Get("topic") == "" {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Topic not provided in URL",
			})
			return
		}

		//get MongoDB client and questions collection
		DB := services.GetConnection()
		questionsCollection := services.GetCollection(DB, "questions")

		//retrieve topic from param
		topic := params.Get("topic")

		//MongoDB filter to search the question related to the `topic` param
		filter := bson.M{"$text": bson.M{"$search": topic}}
		cursor, err := questionsCollection.Find(context, filter)
		if err != nil {
			return
		}
		var questions []models.Question
		if err = cursor.All(context, &questions); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, questions)
	}
}

func UploadQuestionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		DB := services.GetConnection()
		questionsCollection := services.GetCollection(DB, "questions")
		topicsCollection := services.GetCollection(DB, "topics")

		file, _ := c.FormFile("questions_file")
		topic := c.PostForm("topic")
		if topic == "" || len(topic) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Topic not provided",
			})
			return
		}

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

		//for i := 0; i < len(questions); i++ {
		//	fmt.Printf("%v\n", questions[i])
		//}
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

		topicRes, err := topicsCollection.InsertOne(c, topic)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Server error. Please try again later",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"insertedIDs":   fmt.Sprintf("%v", res.InsertedIDs),
			"insertedTopic": fmt.Sprintf("%v", topicRes.InsertedID),
		})

	}
}
