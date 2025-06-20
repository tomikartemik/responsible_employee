package service

import (
	"responsible_employee/internal/model"
	"responsible_employee/internal/repository"
	"responsible_employee/internal/utils"
)

type QuestionService struct {
	repo     repository.Question
	repoUser repository.User
}

func NewQuestionService(repo repository.Question, repoUser repository.User) *QuestionService {
	return &QuestionService{repo: repo, repoUser: repoUser}
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

func (s *QuestionService) CheckUserAnswers(userID string, answers model.TestInput) (int, error) {
	points := 0

	user, err := s.repoUser.GetUserByID(userID)
	if err != nil {
		return err
	}

	for _, userAnswer := range answers.UserAnswers {
		questionOutput, err := s.repo.QuestionByID(userAnswer.QuestionID)
		if err != nil {
			return err
		}

		var selectedAnswer *model.AnswerOption
		for _, answer := range questionOutput.Answers {
			if answer.ID == userAnswer.AnswerID {
				selectedAnswer = &answer
				break
			}
		}

		if selectedAnswer != nil && selectedAnswer.IsCorrect {
			points += 10
		}
	}

	return s.repoUser.UpdateUserPoints(utils.AddPoints(user, points))
}
