package main

import (
	groupService "./usecase/group"
	userService "./usecase/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, userService.GetUsers())
	})
	r.GET("/users/:userId", func(c *gin.Context) {
		userId := c.Param("userId")
		c.JSON(http.StatusOK, userService.GetUser(userId))
	})
	r.GET("/groups", func(c *gin.Context) {
		c.JSON(http.StatusOK, groupService.GetGroups())
	})
	r.Run()
}
