package service

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"responsible_employee/internal/model"
	"responsible_employee/internal/repository"
	"responsible_employee/internal/utils"
)

type UserService struct {
	repo     repository.User
	repoTask repository.Task
}

func NewUserService(repo repository.User, repoTask repository.Task) *UserService {
	return &UserService{repo: repo, repoTask: repoTask}
}

func (s *UserService) SignUp(userData model.User) error {
	userData.Password = utils.GeneratePasswordHash(userData.Password)
	userData.ID = uuid.Must(uuid.NewV4()).String()
	userData.Points = 0
	userData.MaxPoints = 0
	userData.Rank = 0
	return s.repo.SignUp(userData)
}

func (s *UserService) SignIn(userData model.SignInInput) (model.SignInOutput, error) {
	userData.Password = utils.GeneratePasswordHash(userData.Password)
	fmt.Println("service sign in " + userData.Password)

	userID, err := s.repo.SignIn(userData)
	if err != nil {
		return model.SignInOutput{}, err
	}

	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return model.SignInOutput{}, err
	}

	token, err := CreateToken(user.ID)
	if err != nil {
		return model.SignInOutput{}, err
	}

	userOutput := model.SignInOutput{
		Token: token,
		User:  utils.UserToUserOutput(user),
	}

	return userOutput, nil
}

func (s *UserService) ChangePassword(userID string, password, newPassword string) error {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	if user.Password != utils.GeneratePasswordHash(password) {
		return errors.New("password incorrect")
	}

	err = s.repo.ChangePassword(userID, utils.GeneratePasswordHash(newPassword))
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUsersSortedByPoints() ([]model.UserInfoTable, error) {
	tableInfo := []model.UserInfoTable{}
	users, err := s.repo.GetUsersSortedByPoints()

	if err != nil {
		return tableInfo, err
	}

	for _, user := range users {
		tableInfo = append(tableInfo, model.UserInfoTable{ID: user.ID, Login: user.Login, Image: user.Image, Points: user.Points})
	}

	return tableInfo, nil
}

func (s *UserService) CompleteTask(userID, taskID string) error {
	user, err := s.repo.GetUserByID(userID)

	fmt.Println(userID, taskID)

	if err != nil {
		logrus.Error(err)
		return err
	}

	task, err := s.repoTask.TaskByID(taskID)

	if err != nil {
		logrus.Error(err)
		return err
	}

	err = s.repoTask.CompleteTask(taskID)

	if err != nil {
		logrus.Error(err)
		return err
	}

	err = s.repo.UpdateUserPoints(userID, user.Points+task.Points)

	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
