package service

import (
	"mime/multipart"
	"responsible_employee/internal/model"
	"responsible_employee/internal/repository"
)

type Service struct {
	User
	Task
	Report
	Violation
	Photo
	Message
	Question
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:      NewUserService(repos.User, repos.Task, repos.Violation),
		Task:      NewTaskService(repos.Task, repos.Violation, repos.User, repos.Message),
		Report:    NewReportService(repos.Report, repos.Task, repos.User, repos.Message),
		Violation: NewViolationService(repos.Violation),
		Photo:     NewPhotoService(repos.Task, repos.Report),
		Message:   NewMessageService(repos.Message),
		Question:  NewQuestionService(repos.Question, repos.User),
	}
}

type User interface {
	SignUp(userData model.User) error
	SignIn(userData model.SignInInput) (model.SignInOutput, error)
	GetUserByID(userID string) (model.UserOutput, error)
	ChangePassword(userID string, password, newPassword string) error
	GetUsersSortedByPoints() ([]model.UserInfoTable, error)
	TakeTask(userID, taskID string) error
	ChangePasswordByMail(username string) error
}

type Task interface {
	CreateTask(task model.Task, reportedUserID string) (string, error)
	GetAllTasksInfo() ([]model.TasksShortInfo, error)
	GetAllTasksForAnalise() ([]model.TaskForAnalise, error)
	TaskByID(taskID string) (model.Task, error)
}

type Report interface {
	RegisterReport(report model.Report) error
	ReportByID(reportID string) (model.Report, error)
}

type Violation interface {
	GetAllViolations() ([]model.Violation, error)
	GetViolationByCategory(category string) ([]model.Violation, error)
	GetViolationByID(id int) (model.Violation, error)
}

type Photo interface {
	SaveTaskPhoto(taskID string, photo *multipart.FileHeader) error
	SaveReportPhoto(reportID string, photo *multipart.FileHeader) error
}

type Message interface {
	MessagesByUserID(userID string) ([]model.Message, error)
	ReadMessage(messageID int) error
}

type Question interface {
	QuestionByID(questionID int) (model.QuestionOutput, error)
	GenerateTest() ([]model.QuestionOutput, error)
	CheckUserAnswers(userID string, answers model.TestInput) (model.TestResult, error)
}
