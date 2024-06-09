package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zeekhoks/test-api-backend/models"
	"github.com/zeekhoks/test-api-backend/services"
	"log"
	"net/http"
)

func CreateNewUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User

		if err := ctx.Bind(&user); err != nil {
			log.Println("Failed to bind incoming payload with Gin", err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userExists, _ := services.UserExists(user.Username)
		if userExists {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
			return
		}

		createdUser, err := services.CreateUser(user)
		if err != nil {
			log.Println("Failed to create user", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unable to create an user"})
			return
		}

		ctx.JSON(http.StatusCreated, createdUser)
	}
}
