package controllers

import (
	"net/http"
	"strconv"
	"task-api/models"

	"github.com/gin-gonic/gin"
)
var taskMap = make(map[int]models.Task)
var nextID = 1

// GetTasks godoc
// @Summary      Get all tasks
// @Description  Returns all tasks
// @Tags         tasks
// @Produce      json
// @Success      200  {array}  models.Task
// @Router       /tasks [get]
func GetTasks(c *gin.Context) {
	tasks := make([]models.Task, 0, len(taskMap))
	for _, task := range taskMap {
		tasks = append(tasks, task)
	}
	c.JSON(http.StatusOK, tasks)
}

// CreateTask godoc
// @Summary      Create a new task
// @Description  Adds a new task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task  body  models.Task  true  "Task to create"
// @Success      201   {object}  models.Task
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /tasks [post]
func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	newTask.ID = nextID
	nextID++
	taskMap[newTask.ID] = newTask
	c.JSON(http.StatusCreated, newTask)
}

// UpdateTask godoc
// @Summary      Update an existing task
// @Description  Updates values on a task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "Task ID"
// @Param        task  body      models.Task true  "Task for update"
// @Success      200   {object}  models.Task
// @Failure      400   {object}  models.ErrorResponse
// @Failure      404   {object}  models.ErrorResponse
// @Router       /tasks/{id} [put]
func UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid ID"})
		return
	}

	existingTask, exists := taskMap[id]
	if !exists {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Task not found"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	existingTask.Title = updatedTask.Title
	existingTask.Completed = updatedTask.Completed
	taskMap[id] = existingTask

	c.JSON(http.StatusOK, existingTask)
}

// DeleteTask godoc
// @Summary      Delete a task
// @Description  Deletes a task by its ID
// @Tags         tasks
// @Param        id   path      int  true  "Task ID"
// @Success      204  {object}  nil
// @Failure      400  {object}  map[string]string  "Invalid ID"
// @Failure      404  {object}  map[string]string  "Task not found"
// @Router       /tasks/{id} [delete]
func DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid ID"})
		return
	}

	if _, exists := taskMap[id]; !exists {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Task not found"})
		return
	}

	delete(taskMap, id)
	c.Status(http.StatusNoContent)
}