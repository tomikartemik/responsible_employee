package repository

import (
	"gorm.io/gorm"
	"responsible_employee/internal/model"
)

type ViolationRepository struct {
	db *gorm.DB
}

func NewViolationRepository(db *gorm.DB) *ViolationRepository {
	return &ViolationRepository{db: db}
}

func (r *ViolationRepository) GetAllViolations() ([]model.Violation, error) {
	var tasks []model.Violation

	err := r.db.Find(&tasks).Error
	if err != nil {
		return []model.Violation{}, err
	}

	return tasks, nil
}

func (r *ViolationRepository) GetViolationByCategory(category string) ([]model.Violation, error) {
	var tasks []model.Violation

	err := r.db.Where("category = ?", category).Find(&tasks).Error
	if err != nil {
		return []model.Violation{}, err
	}

	return tasks, nil
}

func (r *ViolationRepository) GetViolationByID(id int) (model.Violation, error) {
	var task model.Violation

	err := r.db.Where("id = ?", id).First(&task).Error
	if err != nil {
		return model.Violation{}, err
	}

	return task, nil
}
