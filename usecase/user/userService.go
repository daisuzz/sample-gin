package user

import "github.com/gin-gonic/gin"

func GetUsers() gin.H {
	return gin.H{
		"name": "pong",
	}
}

func GetUser(userId string) gin.H {
	return gin.H{
		"userId": "1",
		"name":   "Ichiro",
	}
}
