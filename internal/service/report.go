package service

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"responsible_employee/internal/model"
	"responsible_employee/internal/repository"
	"responsible_employee/internal/utils"
	"time"
)

type ReportService struct {
	repo        repository.Report
	repoTask    repository.Task
	repoUser    repository.User
	repoMessage repository.Message
    repoPointEvent repository.PointEvent
}

func NewReportService(repo repository.Report, repoTask repository.Task, repoUser repository.User, repoMessage repository.Message, repoPointEvent repository.PointEvent) *ReportService {
    return &ReportService{repo: repo, repoTask: repoTask, repoUser: repoUser, repoMessage: repoMessage, repoPointEvent: repoPointEvent}
}

func (s *ReportService) ReportByID(reportID string) (model.Report, error) {
	report, err := s.repo.ReportByID(reportID)

	if err != nil {
		return model.Report{}, err
	}

	return report, nil
}

func (s *ReportService) RegisterReport(report model.Report) error {
	task, err := s.repoTask.TaskByID(report.TaskID)
	if err != nil {
		return err
	}

	if time.Now().After(task.EndDate.Add(48 * time.Hour)) {
		return errors.New("Срок выполнения задания истек!")
	}

	report.ID = uuid.Must(uuid.NewV4()).String()
	report.ReportedAt = time.Now()

	err = s.repo.CreateReport(report)

	if err != nil {
		return err
	}

	task.ReportID = report.ID
	task.Status = "Completed"
	task.Points = 100 - (s.hoursSincePublication(task.DateReported) * 2)

	err = s.repoTask.UpdateTask(task)
	if err != nil {
		return err
	}

	user, err := s.repoUser.GetUserByID(report.UserID)

	if err != nil {
		return err
	}

	err = s.repoMessage.CreateMessage(model.Message{
		UserID: task.ResponsiblePersonID,
		Text:   fmt.Sprintf("Вы успешно устранили нарушение и заработали %d баллов!", task.Points),
	})

	if err != nil {
		return err
	}

    user = utils.AddPoints(user, task.Points)
    err = s.repoUser.UpdateUserPoints(user)

	if err != nil {
		return err
	}

    // Логируем событие начисления баллов за выполнение
    _ = s.repoPointEvent.Create(model.PointEvent{UserID: task.ResponsiblePersonID, Source: utils.PointsSourceCompleteTask, Points: task.Points, TaskID: &task.ID, ReportID: &report.ID})

	return nil
}

func (s *ReportService) hoursSincePublication(publishedTime time.Time) int {
	now := time.Now()
	duration := now.Sub(publishedTime)
	return int(duration.Hours())
}
