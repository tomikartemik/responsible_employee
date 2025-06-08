package repository

import (
	"gorm.io/gorm"
	"responsible_employee/internal/model"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) CreateMessage(message model.Message) error {
	return r.db.Create(&message).Error
}

func (r *MessageRepository) MessagesByUserID(userID string) ([]model.Message, error) {
	var messages []model.Message

	if err := r.db.Where("user_id = ?", userID).Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *MessageRepository) ReadMessage(messageID int) error {
	var message model.Message
	if err := r.db.First(&message, messageID).Error; err != nil {
		return err
	}

	message.Read = true

	return r.db.Save(&message).Error
}
