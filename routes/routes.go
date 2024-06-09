package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zeekhoks/test-api-backend/controllers"
	"github.com/zeekhoks/test-api-backend/middleware"
)

func GetRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/questions", middleware.BasicAuth(), controllers.GetDisplayQuestionsByTopicHandler())
	router.POST("/questions", middleware.AdminAuth(), controllers.UploadQuestionHandler())
	router.GET("/topics", middleware.BasicAuth(), controllers.GetAllTopics())
	router.POST("/user", controllers.CreateNewUser())

	return router
}
