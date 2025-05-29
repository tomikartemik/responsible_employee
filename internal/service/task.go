package service

import (
	"github.com/gofrs/uuid"
	"responsible_employee/internal/model"
	"responsible_employee/internal/repository"
	"responsible_employee/internal/utils"
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

func (s *TaskService) GetAllTasks() ([]model.TasksShortInfo, error) {
	tasks, err := s.repo.GetAllTasks()

	if err != nil {
		return []model.TasksShortInfo{}, err
	}

	taskShortInfo := make([]model.TasksShortInfo, 0)
	for _, task := range tasks {
		taskShortInfo = append(taskShortInfo, utils.TaskToTaskShortInfo(task))
	}

	return taskShortInfo, nil
}

func (s *TaskService) TaskByID(taskID string) (model.Task, error) {
	task, err := s.repo.TaskByID(taskID)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}
