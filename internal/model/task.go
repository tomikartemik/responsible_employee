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
	Points              int       `json:"points"`
	Status              string    `gorm:"not null" json:"status"`
	ReportedUserId      string    `json:"reportedUserId"`
	ResponsiblePersonID string    `gorm:"foreignKey:ID;references:ResponsiblePersonID" json:"responsiblePerson"`
	ReportID            string    `json:"reportId"`
	TimeLeft            string    `gorm:"not null" json:"timeLeft"`
}

type TasksShortInfo struct {
	ID          string    `gorm:"primaryKey;unique" json:"id"`
	ViolationID int       `gorm:"not null" json:"violation_Id"`
	Violation   Violation `gorm:"foreignKey:ViolationID" json:"violation"`
	Points      int       `json:"points"`
	TimeLeft    string    `gorm:"not null" json:"timeLeft"`
}
