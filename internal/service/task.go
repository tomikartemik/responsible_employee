package service

import (
	"fmt"
	"github.com/gofrs/uuid"
	"responsible_employee/internal/model"
	"responsible_employee/internal/repository"
	"responsible_employee/internal/utils"
	"time"
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
	task.DateReported = time.Now()
	task.EndDate = task.DateReported.Add(48 * time.Hour)
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
		Text:   fmt.Sprintf("Вы успешно зарегистрировали нарушение и заработали %d баллов!", points),
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

func (s *TaskService) GetAllTasksInfo() ([]model.TasksShortInfo, error) {
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

func (s *TaskService) GetAllTasksForAnalise() ([]model.TaskForAnalise, error) {
	tasks, err := s.repo.GetAllTasksForAnalise()
	if err != nil {
		return []model.TaskForAnalise{}, err
	}

	tasksForAnalise := make([]model.TaskForAnalise, 0)
	for _, task := range tasks {

		reportedUser, err := s.repoUser.GetUserByID(task.ReportedUserId)
		if err != nil {
			return nil, err
		}

		responsiblePerson, err := s.repoUser.GetUserByID(task.ResponsiblePersonID)
		if err != nil {
			return nil, err
		}

		taskForAnalise := model.TaskForAnalise{
			ID:                task.ID,
			Violation:         task.Violation,
			Description:       task.Description,
			Suggestion:        task.Suggestion,
			ImageUrl:          task.ImageUrl,
			DateReported:      task.DateReported,
			Status:            task.Status,
			ReportedUser:      reportedUser.FullName,
			ResponsiblePerson: responsiblePerson.FullName,
		}

		tasksForAnalise = append(tasksForAnalise, taskForAnalise)
	}
	return tasksForAnalise, nil
}

func (s *TaskService) TaskByID(taskID string) (model.Task, error) {
	task, err := s.repo.TaskByID(taskID)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}
