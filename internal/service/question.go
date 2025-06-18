package service

import (
	"responsible_employee/internal/model"
	"responsible_employee/internal/repository"
)

type QuestionService struct {
	repo repository.Question
}

func NewQuestionService(repo repository.Question) *QuestionService {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) QuestionByID(questionID int) (model.QuestionOutput, error) {
	question, err := s.repo.QuestionByID(questionID)

	if err != nil {
		return model.QuestionOutput{}, err
	}

	return question, nil
}
