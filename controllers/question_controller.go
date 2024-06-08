package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zeekhoks/test-api-backend/services"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

var DB *mongo.Client = services.GetConnection()

//var questionCollection = services.GetCollection(DB, "questions")

func GetDisplayQuestionsByTopicHandler() gin.HandlerFunc {

	return func(context *gin.Context) {

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

//func UploadQuestionHandler() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//		file, _ := c.FormFile("questions_file")
//		f, _ := file.Open()
//		defer f.Close()
//		content := make([]byte, file.Size)
//		f.Read(content)
//		err := json.Unmarshal(byteValue, &questionCollection)
//		var question models.Question
//		defer cancel()
//		newQuestion := models.Question{
//			ID:            primitive.NewObjectID(),
//			QuestionName:  question.QuestionName,
//			Options:       question.Options,
//			CorrectAnswer: question.CorrectAnswer,
//			Distractors:   question.Distractors,
//		}
//		result, err := questionCollection.InsertOne(ctx, newQuestion)
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
//			return
//		}
//		log.Println(string(content))
//		context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
//		c.JSON(http.StatusCreated)
//	}
//}
