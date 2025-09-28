package model

type Question struct {
	ID       int    `gorm:"primaryKey;uniqueIndex;autoIncrement" json:"id"`
	Text     string `gorm:"not null" json:"text"`
	Category string `gorm:"not null;default:'general'" json:"category"`
}

type AnswerOption struct {
	ID         int    `gorm:"primaryKey;uniqueIndex;autoIncrement" json:"id"`
	QuestionID int    `gorm:"not null" json:"question_id"`
	Label      string `gorm:"type:char(1);not null" json:"label"`
	Text       string `gorm:"not null" json:"text"`
	IsCorrect  bool   `gorm:"not null" json:"is_correct"`
	Category   string `gorm:"not null;default:'general'" json:"category"`
}

type QuestionOutput struct {
	Question Question       `json:"question"`
	Answers  []AnswerOption `json:"answers"`
}

type UserAnswer struct {
	QuestionID int `json:"question_id"`
	AnswerID   int `json:"answer_id"`
}

type TestInput struct {
	UserAnswers []UserAnswer `json:"user_answers"`
}

type WrongAnswer struct {
	Question       Question       `json:"question"`
	UserAnswer     AnswerOption   `json:"user_answer"`
	CorrectAnswer  AnswerOption   `json:"correct_answer"`
	AllAnswers     []AnswerOption `json:"all_answers"`
}

type TestResult struct {
	Points        int            `json:"points"`
	WrongAnswers  []WrongAnswer  `json:"wrong_answers"`
	Message       string         `json:"message"`
}
