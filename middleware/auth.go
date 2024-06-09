package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zeekhoks/test-api-backend/services"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func BasicAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		username, password, ok := context.Request.BasicAuth()
		log.Println("Authenticating user", bson.M{"user username": username})
		if !ok {
			log.Println("Unable to authenticate user, failed to parse auth string")
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid base64 encoding"})
		} else {
			user, err := services.GetUserByUsername(username)
			if err != nil {
				log.Println("Unable to get user with username", err)
				context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unable to find username"})
			} else {
				success := services.CheckPasswordHash(password, user.Password)
				if success {
					context.Set("loggedInAccount", user)
					context.Next()
				} else {
					log.Println("Authentication failed")
					context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization failed"})
				}
			}
		}
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		username, password, ok := context.Request.BasicAuth()
		log.Println("Authenticating admin", bson.M{"admin username": username})
		if !ok {
			log.Println("Unable to authenticate admin, failed to parse auth string")
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid base64 encoding"})
		} else {
			user, err := services.GetUserByUsername(username)
			if err != nil {
				log.Println("Unable to get user with username", err)
				context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unable to find username"})
			} else {
				success := services.CheckPasswordHash(password, user.Password)
				if success {
					if user.IsAdmin {
						context.Set("loggedInAccount", user)
						context.Next()
					} else {
						log.Println("Authorization failed")
						context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization failed"})
					}
				} else {
					log.Println("Authentication failed")
					context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
				}
			}
		}
	}
}
