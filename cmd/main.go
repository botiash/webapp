package main

import (
	"botiash/webapp/internal/handlers"
	"botiash/webapp/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	taskRepo := &repository.InMemoryTaskRepository{}
	taskHandler := &handlers.TaskHandler{
		Repo: taskRepo,
	}

	r.GET("/tasks", taskHandler.GetTasks)
	r.POST("/tasks", taskHandler.CreateTask)

	r.Run(":8080")
}
