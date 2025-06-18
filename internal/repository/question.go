package repository

import (
	"gorm.io/gorm"
	"responsible_employee/internal/model"
)

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (r *QuestionRepository) QuestionByID(questionID int) (model.QuestionOutput, error) {
	var question model.Question
	err := r.db.Where("id = ?", questionID).First(&question).Error
	if err != nil {
		return model.QuestionOutput{}, err
	}

	var answers []model.AnswerOption
	err = r.db.Where("question_id = ?", questionID).Find(&answers).Error
	if err != nil {
		return model.QuestionOutput{}, err
	}

	return model.QuestionOutput{
		Question: question,
		Answers:  answers,
	}, nil
}
