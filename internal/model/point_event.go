package model

type UserPoints struct {
	UserID         string `json:"userId"`
	FullName       string `json:"fullName"`
	MonthlyTotal   int    `json:"monthlyTotal"`
	YearlyTotal    int    `json:"yearlyTotal"`
	LastMonthTotal int    `json:"lastMonthTotal"`
	LastYearTotal  int    `json:"lastYearTotal"`
}
