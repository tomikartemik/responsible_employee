package model

type RiskReport struct {
	ID          string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Category    string `gorm:"not null" json:"category"`
	Type        string `gorm:"not null" json:"type"`
	Description string `gorm:"not null" json:"description"`
	Suggestion  string `gorm:"not null" json:"suggestion"`
	ImageUrl    string `json:"imageUrl"`
}
