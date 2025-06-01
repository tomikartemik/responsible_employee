package service

import (
	"responsible_employee/internal/model"
	"responsible_employee/internal/repository"
)

type Service struct {
	User
	Task
	Report
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:   NewUserService(repos.User, repos.Task),
		Task:   NewTaskService(repos.Task),
		Report: NewReportService(repos.Report, repos.Task, repos.User),
	}
}

type User interface {
	SignUp(userData model.User) error
	SignIn(userData model.SignInInput) (model.SignInOutput, error)
	ChangePassword(userID string, password, newPassword string) error
	GetUsersSortedByPoints() ([]model.UserInfoTable, error)
	TakeTask(userID, taskID string) error
}

type Task interface {
	CreateTask(task model.Task) error
	GetAllTasks() ([]model.TasksShortInfo, error)
	TaskByID(taskID string) (model.Task, error)
}

type Report interface {
	RegisterReport(report model.Report) error
}
