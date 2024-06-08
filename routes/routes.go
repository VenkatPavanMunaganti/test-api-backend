package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zeekhoks/test-api-backend/controllers"
)

func GetRouter() *gin.Engine {

	router := gin.Default()

	router.POST("/questions", controllers.GetAddQuestionsHandler())
	router.GET("/questions", controllers.GetDisplayQuestionsByTopicHandler())

	return router

}
