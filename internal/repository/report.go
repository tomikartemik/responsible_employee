package repository

import (
	"fmt"
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

func (r *ReportRepository) ReportByID(reportID string) (model.Report, error) {
	var report model.Report

	err := r.db.Where("id = ?", reportID).First(&report).Error
	if err != nil {
		return model.Report{}, err
	}

	return report, nil
}

func (r *ReportRepository) AddPhotoToReport(reportID, photoUrl string) error {
	fmt.Println(photoUrl)
	return r.db.Model(model.Report{}).Where("task_id = ?", reportID).Update("image_url", photoUrl).Error
}

func (r *ReportRepository) UpdateReport(report model.Report) error {
	return r.db.Save(&report).Error
}
