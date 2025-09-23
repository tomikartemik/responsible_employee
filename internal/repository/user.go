package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"responsible_employee/internal/model"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) SignUp(user model.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) SignIn(userData model.SignInInput) (string, error) {
	var user model.User

	if err := r.db.Where("login = ?", userData.Login).First(&user).Error; err != nil {
		return "", errors.New("Пользователя с таким никнеймом не существует!")
	}

	fmt.Println("repo sign in " + userData.Password)

	if user.Password != userData.Password {
		return "", errors.New("Неверный пароль!")
	}

	return user.ID, nil
}

func (r *UserRepository) GetUserByID(userID string) (model.User, error) {
	var user model.User
	currentTime := time.Now()

	err := r.db.
		Where("id = ?", userID).
		Preload("Tasks", "status = ? AND end_date > ?", "Taken", currentTime).
		Preload("Tasks.Violation").
		Preload("MyTasks").
		First(&user).
		Error

	if err != nil {
		return model.User{}, errors.New("Пользователь с таким ID не найден!")
	}

	return user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	if err := r.db.Where("login = ?", username).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepository) ChangePassword(userID string, password string) error {
	return r.db.Model(&model.User{}).Where("id = ?", userID).Update("password", password).Error
}

func (r *UserRepository) GetUsersSortedByPoints() ([]model.User, error) {
	var users []model.User
	if err := r.db.Order("monthly_points DESC").Find(&users).Error; err != nil {
		return nil, errors.New("Не удалось получить список пользователей")
	}

	return users, nil
}

func (r *UserRepository) UpdateUserPoints(user model.User) error {
	return r.db.Model(&model.User{}).
		Where("id = ?", user.ID).
		Update("monthly_points", user.MonthlyPoints).
		Update("yearly_points", user.YearlyPoints).
		Update("max_monthly_points", user.MaxMonthlyPoints).
		Update("max_yearly_points", user.MaxYearlyPoints).
		Update("last_month_points", user.LastMonthPoints).
		Update("last_year_points", user.LastYearPoints).
		Update("rank", user.Rank).
		Error
}
