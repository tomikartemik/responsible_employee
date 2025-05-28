package utils

import "responsible_employee/internal/model"

func UserToUserOutput(user model.User) model.UserOutput {
	return model.UserOutput{
		ID:         user.ID,
		Login:      user.Login,
		FullName:   user.FullName,
		Company:    user.Company,
		Department: user.Department,
		Section:    user.Section,
		Position:   user.Position,
		Email:      user.Email,
		Phone:      user.Phone,
		Points:     user.Points,
		Rank:       user.Rank,
	}
}
