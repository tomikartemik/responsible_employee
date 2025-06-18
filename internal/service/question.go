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

func (s *QuestionService) GenerateTest() ([]model.QuestionOutput, error) {
	questionIDs, err := s.repo.RandomQuestionIDs(10)
	if err != nil {
		return nil, err
	}

	var questions []model.QuestionOutput
	for _, id := range questionIDs {
		q, err := s.repo.QuestionByID(id)
		if err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}

	return questions, nil
}
