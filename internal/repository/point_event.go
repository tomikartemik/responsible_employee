package repository

import (
    "time"
    "gorm.io/gorm"
    "responsible_employee/internal/model"
)

type PointEventRepository struct {
    db *gorm.DB
}

func NewPointEventRepository(db *gorm.DB) *PointEventRepository {
    return &PointEventRepository{db: db}
}

func (r *PointEventRepository) Create(event model.PointEvent) error {
    return r.db.Create(&event).Error
}

func (r *PointEventRepository) EventsSince(start time.Time) ([]model.PointEvent, error) {
    var events []model.PointEvent
    if err := r.db.Where("created_at >= ?", start).Find(&events).Error; err != nil {
        return nil, err
    }
    return events, nil
}


