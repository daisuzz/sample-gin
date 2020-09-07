package group

import (
	"github.com/gin-gonic/gin"
)

func GetGroups() gin.H {
	return gin.H{
		"groupId": "1",
		"name":    "staff",
	}
}
