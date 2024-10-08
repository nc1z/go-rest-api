package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func NewTaskHandler(c *gin.Context) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	newTask.ID = xid.New().String()
	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}

func GetTasksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func UpdateTaskHandler(c *gin.Context) {
	id := c.Param("id")
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	index := -1
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			index = i
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}
	tasks[index] = task
	c.JSON(http.StatusOK, task)
}

func DeleteTaskHandler(c *gin.Context) {
	id := c.Param("id")
	index := -1
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			index = 1
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "Task has been deleted",
	})
}
