package controllers

import (
	"net/http"
	"strconv"
	"task-api/models"

	"github.com/gin-gonic/gin"
)
var taskMap = make(map[int]models.Task)
var nextID = 1

func GetTasks(c *gin.Context) {
	tasks := make([]models.Task, 0, len(taskMap))
	for _, task := range taskMap {
		tasks = append(tasks, task)
	}
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask.ID = nextID
	nextID++
	taskMap[newTask.ID] = newTask
	c.JSON(http.StatusCreated, newTask)
}

func UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	existingTask, exists := taskMap[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingTask.Title = updatedTask.Title
	existingTask.Completed = updatedTask.Completed
	taskMap[id] = existingTask

	c.JSON(http.StatusOK, existingTask)
}

func DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if _, exists := taskMap[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	delete(taskMap, id)
	c.Status(http.StatusNoContent)
}