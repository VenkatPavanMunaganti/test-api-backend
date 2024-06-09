package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zeekhoks/test-api-backend/controllers"
)

func GetRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/questions", controllers.GetDisplayQuestionsByTopicHandler())
	router.POST("/questions", controllers.UploadQuestionHandler())
	router.GET("/topics", controllers.GetAllTopics())

	return router

}
