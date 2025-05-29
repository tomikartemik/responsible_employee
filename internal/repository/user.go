package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"responsible_employee/internal/model"
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
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return model.User{}, errors.New("Пользователь с таким ID не найден!")
	}

	return user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepository) ChangePassword(userID string, password string) error {
	return r.db.Model(&model.User{}).Where("id = ?", userID).Update("password", password).Error
}
