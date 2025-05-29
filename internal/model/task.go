package model

import "time"

type Task struct {
	ID           string    `gorm:"primaryKey;unique" json:"id"`
	Violation    string    `gorm:"not null" json:"violation"`
	Description  string    `gorm:"not null" json:"description"`
	Suggestion   string    `gorm:"not null" json:"suggestion"`
	ImageUrl     string    `json:"imageUrl"`
	DateReported time.Time `gorm:"not null" json:"dateReported"`
	Points       int       `gorm:"not null" json:"points"`
	Status       string    `gorm:"not null" json:"status"`
	TimeLeft     string    `gorm:"not null" json:"timeLeft"`
}

type TasksShortInfo struct {
	ID        string `gorm:"primaryKey;unique" json:"id"`
	Violation string `gorm:"not null" json:"violation"`
	Points    int    `gorm:"not null" json:"points"`
	TimeLeft  string `gorm:"not null" json:"timeLeft"`
}
