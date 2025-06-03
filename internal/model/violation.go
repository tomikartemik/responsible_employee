package model

type Violation struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Category    string `gorm:"not null"`
	Description string `gorm:"not null"`
	RiskLevel   string `gorm:"not null"`
	Responsible string `gorm:"not null"`
}
