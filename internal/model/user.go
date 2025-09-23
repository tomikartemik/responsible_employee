package model

type User struct {
	ID               string `gorm:"primaryKey;unique" json:"id"`
	Login            string `gorm:"unique;not null" json:"login"`
	Password         string `gorm:"not null" json:"password"`
	Image            string `json:"image"`
	Company          string `gorm:"not null" json:"company"`
	Department       string `gorm:"not null" json:"department"`
	Section          string `gorm:"not null" json:"section"`
	FullName         string `gorm:"not null" json:"fullName"`
	Position         string `gorm:"not null" json:"position"`
	Email            string `gorm:"not null" json:"email"`
	Phone            string `gorm:"not null" json:"phone"`
	Tasks            []Task `gorm:"foreignKey:ResponsiblePersonID" json:"tasks"`
	MyTasks          []Task `gorm:"foreignKey:ReportedUserId" json:"myTasks"`
	MonthlyPoints    int    `gorm:"not null" json:"monthlyPoints"`
	MaxMonthlyPoints int    `gorm:"not null" json:"maxMonthlyPoints"`
	YearlyPoints     int    `gorm:"not null" json:"yearlyPoints"`
	MaxYearlyPoints  int    `gorm:"not null" json:"maxYearlyPoints"`
	LastMonthPoints  int    `gorm:"not null" json:"lastMonthPoints"`
	LastYearPoints   int    `gorm:"not null" json:"lastYearPoints"`
	Rank             int    `json:"rank"`
}

type SignInInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignInOutput struct {
	Token string     `json:"token"`
	User  UserOutput `json:"user"`
}

type ChangePasswordInput struct {
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

type UserOutput struct {
	ID               string `json:"id"`
	Login            string `json:"login"`
	Image            string `json:"image"`
	Company          string `json:"company"`
	Department       string `json:"department"`
	Section          string `json:"section"`
	FullName         string `json:"fullName"`
	Position         string `json:"position"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Tasks            []Task `json:"tasks"`
	MyTasks          []Task `json:"myTasks"`
	MonthlyPoints    int    `json:"monthlyPoints"`
	MaxMonthlyPoints int    `json:"maxMonthlyPoints"`
	YearlyPoints     int    `json:"yearlyPoints"`
	MaxYearlyPoints  int    `json:"maxYearlyPoints"`
	LastMonthPoints  int    `json:"lastMonthPoints"`
	LastYearPoints   int    `json:"lastYearPoints"`
	Rank             int    `json:"rank"`
}

type UserInfoTable struct {
	ID            string `json:"id"`
	Login         string `json:"login"`
	Image         string `json:"image"`
	MonthlyPoints int    `json:"monthlyPoints"`
	YearlyPoints  int    `json:"yearlyPoints"`
	LastMonthPoints int  `json:"lastMonthPoints"`
	LastYearPoints  int  `json:"lastYearPoints"`
}
