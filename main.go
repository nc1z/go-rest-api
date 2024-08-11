package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Task Management API with Golang and Gin"})
}

func main() {
	router := gin.Default()
	router.GET("/", HomeHandler)
	router.GET("/tasks", GetTasksHandler)
	router.POST("/tasks", NewTaskHandler)
	router.PUT("/tasks/:id", UpdateTaskHandler)
	router.DELETE("/tasks/:id", DeleteTaskHandler)
	err := router.Run()
	if err != nil {
		return
	}
}
