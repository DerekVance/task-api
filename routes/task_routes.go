package routes

import (
	"github.com/gin-gonic/gin"
	"task-api/controllers"
)

func RegisterTaskRoutes(r *gin.Engine) {
	taskGroup := r.Group("/tasks")
	{
		taskGroup.GET("/", controllers.GetTasks)
		taskGroup.POST("/", controllers.CreateTask)
		taskGroup.PUT("/:id", controllers.UpdateTask)
	}
}