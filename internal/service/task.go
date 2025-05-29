package service

import (
	"github.com/gofrs/uuid"
	"responsible_employee/internal/model"
	"responsible_employee/internal/repository"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task model.Task) error {
	task.ID = uuid.Must(uuid.NewV4()).String()
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]model.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) TaskByID(taskID string) (model.Task, error) {
	task, err := s.repo.TaskByID(taskID)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}
