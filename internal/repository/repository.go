package repository

import (
	"gorm.io/gorm"
	"responsible_employee/internal/model"
)

type Repository struct {
	User
	Task
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
		Task: NewTaskRepository(db),
	}
}

type User interface {
	SignUp(user model.User) error
	SignIn(userData model.SignInInput) (string, error)
	GetUserByID(userID string) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
	ChangePassword(userID string, password string) error
}

type Task interface {
	CreateTask(task model.Task) error
	GetAllTasks() ([]model.Task, error)
	TaskByID(taskID string) (model.Task, error)
}
