package service

import (
	"responsible_employee/internal/model"
	"responsible_employee/internal/repository"
)

type MessageService struct {
	repo repository.Message
}

func NewMessageService(repo repository.Message) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) MessagesByUserID(userID string) ([]model.Message, error) {
	return s.repo.MessagesByUserID(userID)
}

func (s *MessageService) ReadMessage(messageID int) error {
	return s.repo.ReadMessage(messageID)
}
