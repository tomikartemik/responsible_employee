package utils

import (
	"fmt"
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

	var overdueTasks []model.Task
	err := tc.db.
		Where("end_date < ? AND status NOT IN ?", now, []string{"Completed", "Passed to superiors"}).
		Find(&overdueTasks).Error
	if err != nil {
		log.Printf("Error fetching overdue tasks: %v", err)
		return
	}

	if len(overdueTasks) == 0 {
		return
	}

	for _, task := range overdueTasks {
		if err := tc.db.Model(&model.Task{}).Where("id = ?", task.ID).Update("status", "Passed to superiors").Error; err != nil {
			log.Printf("Error updating task %s: %v", task.ID, err)
			continue
		}

		if task.ResponsiblePersonID != "" {
			message := model.Message{
				UserID: task.ResponsiblePersonID,
				Text:   fmt.Sprintf("Срок выполнения задания \"%s\" истек, задача передана руководству.", task.Description),
			}
			if err := tc.db.Create(&message).Error; err != nil {
				log.Printf("Error notifying responsible user %s for task %s: %v", task.ResponsiblePersonID, task.ID, err)
			}
		}
	}

	log.Printf("Updated %d tasks to 'Passed to superiors' status", len(overdueTasks))
}
