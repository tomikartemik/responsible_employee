package utils

import "responsible_employee/internal/model"

func AddPoints(user model.User, points int) model.User {
	user.MonthlyPoints = user.MonthlyPoints + points
	user.YearlyPoints = user.YearlyPoints + points
	user.MaxMonthlyPoints = max(user.MonthlyPoints, user.MaxMonthlyPoints)
	user.MaxYearlyPoints = max(user.YearlyPoints, user.MaxYearlyPoints)

	switch {
	case user.MaxYearlyPoints >= 200000:
		user.Rank = 10
	case user.MaxYearlyPoints >= 150000:
		user.Rank = 9
	case user.MaxYearlyPoints >= 100000:
		user.Rank = 8
	case user.MaxYearlyPoints >= 80000:
		user.Rank = 7
	case user.MaxYearlyPoints >= 60000:
		user.Rank = 6
	case user.MaxYearlyPoints >= 45000:
		user.Rank = 5
	case user.MaxYearlyPoints >= 30000:
		user.Rank = 4
	case user.MaxYearlyPoints >= 15000:
		user.Rank = 3
	case user.MaxYearlyPoints >= 5000:
		user.Rank = 2
	default:
		user.Rank = 1
	}

	return user
}
