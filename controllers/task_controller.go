package controllers

import (
	"net/http"
	"task-api/models"

	"github.com/gin-gonic/gin"
)

var tasks = []models.Task{}

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}
