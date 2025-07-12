package main

import (
	"task-api/routes"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "task-api/docs"
)

// @title           Task API
// @version         1.0
// @description     Simple task management API in Go using Gin.
// @host            localhost:8080
// @BasePath        /

func main() {
	r := gin.Default()

	routes.RegisterTaskRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}