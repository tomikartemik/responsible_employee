package model

type Question struct {
	ID   int    `gorm:"primaryKey;uniqueIndex;autoIncrement" json:"id"`
	Text string `gorm:"not null" json:"text"`
}

type AnswerOption struct {
	ID         int    `gorm:"primaryKey;uniqueIndex;autoIncrement" json:"id"`
	QuestionID int    `gorm:"not null" json:"question_id"`
	Label      string `gorm:"type:char(1);not null" json:"label"`
	Text       string `gorm:"not null" json:"text"`
	IsCorrect  bool   `gorm:"not null" json:"is_correct"`
}
