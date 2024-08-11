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
	err := router.Run()
	if err != nil {
		return
	}
}
