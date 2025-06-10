package service

import (
	"fmt"
	"github.com/gofrs/uuid"
	"responsible_employee/internal/model"
	"responsible_employee/internal/repository"
	"time"
)

type ReportService struct {
	repo        repository.Report
	repoTask    repository.Task
	repoUser    repository.User
	repoMessage repository.Message
}

func NewReportService(repo repository.Report, repoTask repository.Task, repoUser repository.User, repoMessage repository.Message) *ReportService {
	return &ReportService{repo: repo, repoTask: repoTask, repoUser: repoUser, repoMessage: repoMessage}
}

func (s *ReportService) RegisterReport(report model.Report) error {
	report.ID = uuid.Must(uuid.NewV4()).String()

	err := s.repo.CreateReport(report)
	if err != nil {
		return err
	}

	task, err := s.repoTask.TaskByID(report.TaskID)

	if err != nil {
		return err
	}

	task.ReportID = report.ID
	task.Status = "Completed"
	task.Points = 100 - ((48 - s.hoursSincePublication(task.DateReported)) * 2)

	err = s.repoTask.UpdateTask(task)
	if err != nil {
		return err
	}

	user, err := s.repoUser.GetUserByID(report.UserID)

	if err != nil {
		return err
	}

	err = s.repoMessage.CreateMessage(model.Message{
		UserID: task.ReportedUserId,
		Text:   fmt.Sprintf("Вы успешно устранили нарушение и заработали %d баллов!", task.Points),
	})

	if err != nil {
		return err
	}

	return s.repoUser.UpdateUserPoints(user.ID,
		user.MonthlyPoints+task.Points,
		user.YearlyPoints+task.Points,
		max(user.MaxMonthlyPoints, user.MonthlyPoints+task.Points),
		max(user.MaxYearlyPoints, user.YearlyPoints+task.Points))
}

func (s *ReportService) hoursSincePublication(publishedTime time.Time) int {
	now := time.Now()
	duration := now.Sub(publishedTime)
	return int(duration.Hours())
}
