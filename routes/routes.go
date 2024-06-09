package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zeekhoks/test-api-backend/controllers"
)

func GetRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/questions", controllers.GetDisplayQuestionsByTopicHandler())
	router.POST("/questions", controllers.UploadQuestionHandler())
	router.POST("/quiz", controllers.GenerateQuizHandler())
	router.POST("/quiz/:id/response", controllers.SubmitAnswerHandler())
	router.GET("/quiz/:id/result", controllers.QuizResultHandler())

	return router

}
