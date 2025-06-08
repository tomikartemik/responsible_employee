package model

type Message struct {
	ID     int    `gorm:"primaryKey;uniqueIndex;autoIncrement" json:"id"`
	UserID string `gorm:"not null" json:"user_id"`
	Text   string `gorm:"not null" json:"text"`
	Read   bool   `gorm:"not null" json:"read"`
}
