package repository

import (
	"gorm.io/gorm"
	"responsible_employee/internal/model"
)

type ReportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) CreateReport(report model.Report) error {
	return r.db.Create(&report).Error
}
