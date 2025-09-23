package utils

import "responsible_employee/internal/model"

func UserToUserOutput(user model.User) model.UserOutput {
	return model.UserOutput{
		ID:               user.ID,
		Login:            user.Login,
		FullName:         user.FullName,
		Company:          user.Company,
		Department:       user.Department,
		Section:          user.Section,
		Position:         user.Position,
		Email:            user.Email,
		Phone:            user.Phone,
		Tasks:            user.Tasks,
		MyTasks:          user.MyTasks,
		MonthlyPoints:    user.MonthlyPoints,
		YearlyPoints:     user.YearlyPoints,
		MaxMonthlyPoints: user.MaxMonthlyPoints,
		MaxYearlyPoints:  user.MaxYearlyPoints,
		LastMonthPoints:  user.LastMonthPoints,
		LastYearPoints:   user.LastYearPoints,
		Rank:             user.Rank,
	}
}

func TaskToTaskShortInfo(task model.Task) model.TasksShortInfo {
	return model.TasksShortInfo{
		ID:          task.ID,
		ViolationID: task.ViolationID,
		Violation:   task.Violation,
		Points:      task.Points,
	}
}

func UserToUserPoints(user model.User) model.UserPoints {
	return model.UserPoints{
		UserID:         user.ID,
		FullName:       user.FullName,
		MonthlyTotal:   user.MonthlyPoints,
		LastMonthTotal: user.LastMonthPoints,
		YearlyTotal:    user.YearlyPoints,
		LastYearTotal:  user.LastYearPoints,
	}
}
