package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zeekhoks/test-api-backend/controllers"
	"github.com/zeekhoks/test-api-backend/middleware"
)

func GetRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/questions", middleware.AdminAuth(), controllers.GetDisplayQuestionsByTopicHandler())
	router.POST("/questions", middleware.AdminAuth(), controllers.UploadQuestionHandler())
	router.GET("/topics", middleware.BasicAuth(), controllers.GetAllTopics())
	router.POST("/user", controllers.CreateNewUser())
	router.POST("/quiz", middleware.BasicAuth(), controllers.GenerateQuizHandler())
	router.POST("/quiz/:id/response", middleware.BasicAuth(), controllers.SubmitAnswerHandler())
	router.GET("/quiz/:id/result", middleware.BasicAuth(), controllers.QuizResultHandler())

	return router
}
