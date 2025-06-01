package service

import (
	"github.com/gofrs/uuid"
	"responsible_employee/internal/model"
	"responsible_employee/internal/repository"
)

type ReportService struct {
	repo     repository.Report
	repoTask repository.Task
	repoUser repository.User
}

func NewReportService(repo repository.Report, repoTask repository.Task, repoUser repository.User) *ReportService {
	return &ReportService{repo: repo, repoTask: repoTask, repoUser: repoUser}
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

	err = s.repoTask.UpdateTask(task)
	if err != nil {
		return err
	}

	user, err := s.repoUser.GetUserByID(report.UserID)

	if err != nil {
		return err
	}

	return s.repoUser.UpdateUserPoints(user.ID,
		user.MonthlyPoints+task.Points,
		user.YearlyPoints+task.Points,
		max(user.MaxMonthlyPoints, user.MonthlyPoints+task.Points),
		max(user.MaxYearlyPoints, user.YearlyPoints+task.Points))
}
