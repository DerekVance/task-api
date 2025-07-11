package main

import (
	"task-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.RegisterTaskRoutes(router)
	router.Run(":8080")
}