package service

import (
	"responsible_employee/internal/model"
	"responsible_employee/internal/repository"
)

type Service struct {
	User
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos),
		Task: NewTaskService(repos),
	}
}

type User interface {
	SignUp(userData model.User) error
	SignIn(userData model.SignInInput) (model.SignInOutput, error)
	ChangePassword(userID string, password, newPassword string) error
	GetUsersSortedByPoints() ([]model.UserInfoTable, error)
	CompleteTask(userID, taskID string) error
}

type Task interface {
	CreateTask(task model.Task) error
	GetAllTasks() ([]model.TasksShortInfo, error)
	TaskByID(taskID string) (model.Task, error)
}
