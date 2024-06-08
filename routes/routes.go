package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zeekhoks/test-api-backend/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRouter(DBCon *mongo.Client) *gin.Engine {

	router := gin.Default()

	router.GET("/questions", controllers.GetDisplayQuestionsByTopicHandler())
	//router.POST("/questions", controllers.UploadQuestionHandler())

	return router

}
