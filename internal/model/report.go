package model

import "time"

type Report struct {
	ID         string    `gorm:"primaryKey;unique" json:"id"`
	Comment    string    `json:"comment"`
	ReportedAt time.Time `json:"reportedAt"`
	TaskID     string    `json:"taskId"`
	UserID     string    `json:"userId"`
	ImageUrl   string    `json:"imageUrl"`
}
