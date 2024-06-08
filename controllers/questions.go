package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAddQuestionsHandler() gin.HandlerFunc {

	return func(context *gin.Context) {

		context.JSON(http.StatusOK, gin.H{
			"text": "Get request completed",
		})
	}
}

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
