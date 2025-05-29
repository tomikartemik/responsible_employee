package model

import "time"

type Notification struct {
	ID      string    `gorm:"primaryKey;unique" json:"id"`
	Title   string    `gorm:"not null" json:"title"`
	Message string    `gorm:"not null" json:"message"`
	Date    time.Time `gorm:"not null" json:"date"`
	IsRead  bool      `gorm:"not null;default:false" json:"isRead"`
}
