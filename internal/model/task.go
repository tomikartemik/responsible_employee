package model

import "time"

type Task struct {
	ID                  string     `gorm:"primaryKey;unique" json:"id"`
	ViolationID         int        `gorm:"not null" json:"violationId"`
	Violation           Violation  `gorm:"foreignKey:ViolationID" json:"violation"`
	Description         string     `gorm:"not null" json:"description"`
	Suggestion          string     `gorm:"not null" json:"suggestion"`
	ImageUrl            string     `json:"imageUrl"`
	DateReported        time.Time  `gorm:"not null" json:"dateReported"`
	EndDate             time.Time  `gorm:"not null" json:"endDate"`
	Points              int        `json:"points"`
	Status              string     `gorm:"not null" json:"status"`
	ReportedUserId      string     `json:"reportedUserId"`
	ResponsiblePersonID string     `json:"responsiblePerson" gorm:"default:null"`
	ReportID            string     `json:"reportId"`
	Latitude            *float64   `json:"latitude,omitempty" gorm:"type:decimal(10,8)"`
	Longitude           *float64   `json:"longitude,omitempty" gorm:"type:decimal(11,8)"`
}

type TaskForAnalise struct {
	ID                string     `gorm:"primaryKey;unique" json:"id"`
	Violation         Violation  `gorm:"foreignKey:ViolationID" json:"violation"`
	Description       string     `gorm:"not null" json:"description"`
	Suggestion        string     `gorm:"not null" json:"suggestion"`
	ImageUrl          string     `json:"imageUrl"`
	DateReported      time.Time  `gorm:"not null" json:"dateReported"`
	Status            string     `gorm:"not null" json:"status"`
	ReportedUser      string     `json:"reportedUserId"`
	ResponsiblePerson string     `json:"responsiblePerson" gorm:"default:null"`
	Latitude          *float64   `json:"latitude,omitempty"`
	Longitude         *float64   `json:"longitude,omitempty"`
}

type TasksShortInfo struct {
	ID          string    `gorm:"primaryKey;unique" json:"id"`
	ViolationID int       `gorm:"not null" json:"violation_Id"`
	Violation   Violation `gorm:"foreignKey:ViolationID" json:"violation"`
	Points      int       `json:"points"`
}

type TaskWithCoordinates struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DateReported time.Time `json:"dateReported"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Violation   Violation `json:"violation"`
	ImageUrl    string    `json:"imageUrl,omitempty"`
}

type MapPoint struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
