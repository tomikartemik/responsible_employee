package service

import (
	"responsible_employee/internal/model"
	"responsible_employee/internal/repository"
	"responsible_employee/internal/utils"
)

type PointsService struct {
	repoUser repository.User
}

func NewPointsService(repoUser repository.User) *PointsService {
	return &PointsService{repoUser: repoUser}
}

func (s *PointsService) Summary() ([]model.UserPoints, error) {
	users, err := s.repoUser.GetAllUsers()
	if err != nil {
		return nil, err
	}

	points := make([]model.UserPoints, len(users))
	for i, user := range users {
		points[i] = utils.UserToUserPoints(user)
	}
	return points, nil
}
