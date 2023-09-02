package repository

import "botiash/webapp/internal/models"

type TaskRepository interface {
	GetAllTasks() []models.Task
	CreateTask(task *models.Task)
}

type InMemoryTaskRepository struct {
	Tasks []models.Task
}

func (r *InMemoryTaskRepository) GetAllTasks() []models.Task {
	return r.Tasks
}

func (r *InMemoryTaskRepository) CreateTask(task *models.Task) {
	r.Tasks = append(r.Tasks, *task)
}
