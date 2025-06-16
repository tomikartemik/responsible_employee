package utils

import (
	"log"
	"responsible_employee/internal/model"
	"time"

	"gorm.io/gorm"
)

type TaskChecker struct {
	db     *gorm.DB
	ticker *time.Ticker
	done   chan bool
}

func NewTaskChecker(db *gorm.DB) *TaskChecker {
	return &TaskChecker{
		db:   db,
		done: make(chan bool),
	}
}

func (tc *TaskChecker) Start() {
	tc.ticker = time.NewTicker(10 * time.Minute)

	go func() {
		for {
			select {
			case <-tc.ticker.C:
				tc.checkTasks()
			case <-tc.done:
				tc.ticker.Stop()
				return
			}
		}
	}()

	log.Println("Task checker service started")
}

func (tc *TaskChecker) Stop() {
	tc.done <- true
	log.Println("Task checker service stopped")
}

func (tc *TaskChecker) checkTasks() {
	now := time.Now()

	result := tc.db.Model(&model.Task{}).
		Where("status != ? AND end_date < ?", "Taken", now).
		Update("status", "Passed to superiors")

	if result.Error != nil {
		log.Printf("Error updating tasks: %v", result.Error)
		return
	}

	if result.RowsAffected > 0 {
		log.Printf("Updated %d tasks to 'Passed to superiors' status", result.RowsAffected)
	}
}
