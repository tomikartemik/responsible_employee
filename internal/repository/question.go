package repository

import (
	"gorm.io/gorm"
	"responsible_employee/internal/model"
	"responsible_employee/internal/utils"
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

	// Очистка и обновление, если нужно
	for i := range answers {
		cleaned := utils.CleanLegalRefs(answers[i].Text)
		if answers[i].Text != cleaned {
			answers[i].Text = cleaned
			r.db.Model(&answers[i]).Update("text", cleaned) // обновляем только одно поле
		}
	}

	return model.QuestionOutput{
		Question: question,
		Answers:  answers,
	}, nil
}

func (r *QuestionRepository) RandomQuestionIDs(limit int, category string) ([]int, error) {
	var ids []int

	query := r.db.Model(&model.Question{})
	if category != "" {
		query = query.Where("category = ?", category)
	}

	err := query.Select("id").
		Order("RANDOM()").
		Limit(limit).
		Pluck("id", &ids).Error

	if err != nil {
		return nil, err
	}

	return ids, nil
}
