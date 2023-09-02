package handlers

import (
	"botiash/webapp/internal/models"
	"net/http"

	"botiash/webapp/internal/repository"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	Repo repository.TaskRepository
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks := h.Repo.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.Repo.CreateTask(&task)
	c.JSON(http.StatusCreated, task)
}
