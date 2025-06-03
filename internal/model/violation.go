package model

type Violation struct {
	ID          int    `gorm:"primaryKey;uniqueIndex;autoIncrement" json:"id"`
	Category    string `gorm:"not null" json:"category"`
	Description string `gorm:"not null" json:"description"`
	RiskLevel   string `gorm:"not null" json:"risk_level"`
	Responsible string `gorm:"not null" json:"responsible"`
}
