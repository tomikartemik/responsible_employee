package service

import (
	"responsible_employee/internal/model"
	"responsible_employee/internal/repository"
)

type ViolationService struct {
	repo repository.Violation
}

func NewViolationService(repo repository.Violation) *ViolationService {
	return &ViolationService{repo: repo}
}

func (s *ViolationService) GetAllViolations() ([]model.Violation, error) {
	return s.repo.GetAllViolations()
}

func (s *ViolationService) GetViolationByCategory(category string) ([]model.Violation, error) {
	return s.repo.GetViolationByCategory(category)
}

func (s *ViolationService) GetViolationByID(id int) (model.Violation, error) {
	return s.repo.GetViolationByID(id)
}
