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

func (r *TaskRepository) CreateTask(task model.Task) error {
	task.Status = "Active"
	return r.db.Create(&task).Error
}

func (r *TaskRepository) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task

	err := r.db.Find(&tasks).Where("status = Active").Preload("Violations").Error
	if err != nil {
		return []model.Task{}, err
	}

	return tasks, nil
}

func (r *TaskRepository) TaskByID(taskID string) (model.Task, error) {
	var task model.Task

	err := r.db.Where("id = ?", taskID).Preload("Violations").First(&task).Error
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (r *TaskRepository) UpdateTask(task model.Task) error {
	return r.db.Save(&task).Error
}
