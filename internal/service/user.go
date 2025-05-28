package service

import (
	"errors"
	"fmt"
	"responsible_employee/internal/model"
	"responsible_employee/internal/repository"
	"responsible_employee/internal/utils"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) SignUp(userData model.User) error {
	userData.Password = utils.GeneratePasswordHash(userData.Password)
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
