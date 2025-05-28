package model

type User struct {
	ID         string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Login      string `gorm:"unique;not null" json:"login"`
	Password   string `gorm:"not null" json:"password"`
	Company    string `gorm:"not null" json:"company"`
	Department string `gorm:"not null" json:"department"`
	Section    string `gorm:"not null" json:"section"`
	FullName   string `gorm:"not null" json:"fullName"`
	Position   string `gorm:"not null" json:"position"`
	Email      string `gorm:"not null" json:"email"`
	Phone      string `gorm:"not null" json:"phone"`
	Points     int    `json:"points"`
	Rank       int    `json:"rank"`
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
	ID         string `json:"id"`
	Login      string `json:"login"`
	Company    string `json:"company"`
	Department string `json:"department"`
	Section    string `json:"section"`
	FullName   string `json:"fullName"`
	Position   string `json:"position"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Points     int    `json:"points"`
	Rank       int    `json:"rank"`
}
