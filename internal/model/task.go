package model

import "time"

type Task struct {
	ID                  string    `gorm:"primaryKey;unique" json:"id"`
	ViolationID         int       `gorm:"not null" json:"violationId"`
	Violation           Violation `gorm:"foreignKey:ViolationID" json:"violation"`
	Description         string    `gorm:"not null" json:"description"`
	Suggestion          string    `gorm:"not null" json:"suggestion"`
	ImageUrl            string    `json:"imageUrl"`
	DateReported        time.Time `gorm:"not null" json:"dateReported"`
	EndDate             time.Time `gorm:"not null" json:"endDate"`
	Points              int       `json:"points"`
	Status              string    `gorm:"not null" json:"status"`
	ReportedUserId      string    `json:"reportedUserId"`
	ResponsiblePersonID string    `json:"responsiblePerson" gorm:"default:null"`
	ReportID            string    `json:"reportId"`
}

type TaskForAnalise struct {
	ID                string    `gorm:"primaryKey;unique" json:"id"`
	Violation         Violation `gorm:"foreignKey:ViolationID" json:"violation"`
	Description       string    `gorm:"not null" json:"description"`
	Suggestion        string    `gorm:"not null" json:"suggestion"`
	ImageUrl          string    `json:"imageUrl"`
	DateReported      time.Time `gorm:"not null" json:"dateReported"`
	Status            string    `gorm:"not null" json:"status"`
	ReportedUser      string    `json:"reportedUserId"`
	ResponsiblePerson string    `json:"responsiblePerson" gorm:"default:null"`
}

type TasksShortInfo struct {
	ID          string    `gorm:"primaryKey;unique" json:"id"`
	ViolationID int       `gorm:"not null" json:"violation_Id"`
	Violation   Violation `gorm:"foreignKey:ViolationID" json:"violation"`
	Points      int       `json:"points"`
}
