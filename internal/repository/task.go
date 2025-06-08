package repository

import (
	"gorm.io/gorm"
	"responsible_employee/internal/model"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(task model.Task) (string, error) {
	task.Status = "Active"
	err := r.db.Create(&task).Error
	if err != nil {
		return "", err
	}
	return task.ID, nil
}

func (r *TaskRepository) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task

	err := r.db.Where("status = ?", "Active").Preload("Violation").Find(&tasks).Error
	if err != nil {
		return []model.Task{}, err
	}

	return tasks, nil
}

func (r *TaskRepository) TaskByID(taskID string) (model.Task, error) {
	var task model.Task

	err := r.db.Where("id = ?", taskID).Preload("Violation").First(&task).Error
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (r *TaskRepository) AddPhotoToTask(taskID, photoUrl string) error {
	return r.db.Model(model.Task{}).Where("id = ?", taskID).Update("image_url", photoUrl).Error
}

func (r *TaskRepository) UpdateTask(task model.Task) error {
	return r.db.Save(&task).Error
}
