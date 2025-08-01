package repository

import (
	"gorm.io/gorm"
	"responsible_employee/internal/model"
)

type Repository struct {
	User
	Task
	Report
	Violation
	Message
	Question
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:      NewUserRepository(db),
		Task:      NewTaskRepository(db),
		Report:    NewReportRepository(db),
		Violation: NewViolationRepository(db),
		Message:   NewMessageRepository(db),
		Question:  NewQuestionRepository(db),
	}
}

type User interface {
	SignUp(user model.User) error
	SignIn(userData model.SignInInput) (string, error)
	GetUserByID(userID string) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
	ChangePassword(userID string, password string) error
	GetUsersSortedByPoints() ([]model.User, error)
	UpdateUserPoints(user model.User) error
}

type Task interface {
	CreateTask(task model.Task) (string, error)
	GetAllTasks() ([]model.Task, error)
	TaskByID(taskID string) (model.Task, error)
	UpdateTask(task model.Task) error
	AddPhotoToTask(taskID, photoUrl string) error
	GetAllTasksForAnalise() ([]model.Task, error)
}

type Report interface {
	CreateReport(report model.Report) error
	ReportByID(reportID string) (model.Report, error)
	UpdateReport(report model.Report) error
}

type Violation interface {
	GetAllViolations() ([]model.Violation, error)
	GetViolationByCategory(category string) ([]model.Violation, error)
	GetViolationByID(id int) (model.Violation, error)
}

type Message interface {
	CreateMessage(message model.Message) error
	MessagesByUserID(userID string) ([]model.Message, error)
	ReadMessage(messageID int) error
}

type Question interface {
	QuestionByID(questionID int) (model.QuestionOutput, error)
	RandomQuestionIDs(limit int) ([]int, error)
}
