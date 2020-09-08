package main

import (
	"./infrastructure"
	"./usecase"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()

	// setup logfile
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	userRepository := infrastructure.CreateUserRepository()
	userService := usecase.CreateUserService(userRepository)

	groupRepository := infrastructure.CreateGroupRepository()
	groupService := usecase.CreateGroupService(groupRepository)

	v1 := r.Group("/v1")
	{
		v1.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, userService.GetUsers())
		})
		v1.GET("/users/:userId", func(c *gin.Context) {
			userId := c.Param("userId")
			c.JSON(http.StatusOK, userService.GetUser(userId))
		})
		v1.GET("/groups", func(c *gin.Context) {
			c.JSON(http.StatusOK, groupService.GetGroups())
		})
	}
	r.Run()
}
