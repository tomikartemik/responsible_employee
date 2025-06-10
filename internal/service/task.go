package service

import (
	"fmt"
	"github.com/gofrs/uuid"
	"responsible_employee/internal/model"
	"responsible_employee/internal/repository"
	"responsible_employee/internal/utils"
)

type TaskService struct {
	repo          repository.Task
	repoViolation repository.Violation
	repoUser      repository.User
	repoMessage   repository.Message
}

func NewTaskService(repo repository.Task, repoViolation repository.Violation, repoUser repository.User, repoMessage repository.Message) *TaskService {
	return &TaskService{repo: repo, repoViolation: repoViolation, repoUser: repoUser, repoMessage: repoMessage}
}

func (s *TaskService) CreateTask(task model.Task, reportedUserID string) (string, error) {
	task.ID = uuid.Must(uuid.NewV4()).String()
	task.ReportedUserId = reportedUserID

	violation, err := s.repoViolation.GetViolationByID(task.ViolationID)
	if err != nil {
		return "", err
	}

	points := 20

	if violation.RiskLevel == "Средний" {
		points = 25
	}

	if violation.RiskLevel == "Высокий" {
		points = 30
	}

	user, err := s.repoUser.GetUserByID(task.ReportedUserId)

	if err != nil {
		return "", err
	}

	user = utils.AddPoints(user, points)
	err = s.repoUser.UpdateUserPoints(user)

	if err != nil {
		return "", err
	}

	err = s.repoMessage.CreateMessage(model.Message{
		UserID: task.ReportedUserId,
		Text:   fmt.Sprintf("Вы успешно зарегестрировали нарушение и заработали %d баллов!", task.Points),
	})

	if err != nil {
		return "", err
	}

	taskID, err := s.repo.CreateTask(task)
	if err != nil {
		return "", err
	}
	return taskID, nil
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
