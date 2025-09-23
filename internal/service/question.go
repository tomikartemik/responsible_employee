package service

import (
	"fmt"
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

func (s *QuestionService) CheckUserAnswers(userID string, answers model.TestInput) (model.TestResult, error) {
	points := 0
	var wrongAnswers []model.WrongAnswer

	user, err := s.repoUser.GetUserByID(userID)
	if err != nil {
		return model.TestResult{}, err
	}

	for _, userAnswer := range answers.UserAnswers {
		questionOutput, err := s.repo.QuestionByID(userAnswer.QuestionID)
		if err != nil {
			return model.TestResult{}, err
		}

		var selectedAnswer *model.AnswerOption
		var correctAnswer *model.AnswerOption

		for _, answer := range questionOutput.Answers {
			if answer.ID == userAnswer.AnswerID {
				selectedAnswer = &answer
			}
			if answer.IsCorrect {
				correctAnswer = &answer
			}
		}

		if selectedAnswer != nil && selectedAnswer.IsCorrect {
			points += 10
		} else {
			// Добавляем неправильный ответ в список
			if selectedAnswer != nil && correctAnswer != nil {
				wrongAnswers = append(wrongAnswers, model.WrongAnswer{
					Question:      questionOutput.Question,
					UserAnswer:    *selectedAnswer,
					CorrectAnswer: *correctAnswer,
					AllAnswers:    questionOutput.Answers,
				})
			}
		}
	}

	// Обновляем баллы пользователя
	updated := utils.AddPoints(user, points)
	err = s.repoUser.UpdateUserPoints(updated)
	if err != nil {
		return model.TestResult{}, err
	}

	message := fmt.Sprintf("Получено %d баллов!", points)
	if len(wrongAnswers) > 0 {
		message += fmt.Sprintf(" Неправильных ответов: %d", len(wrongAnswers))
	}

	return model.TestResult{
		Points:       points,
		WrongAnswers: wrongAnswers,
		Message:      message,
	}, nil
}
