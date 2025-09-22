package repository

import (
    "gorm.io/gorm"
    "responsible_employee/internal/model"
)

type MetaRepository struct {
    db *gorm.DB
}

func NewMetaRepository(db *gorm.DB) *MetaRepository {
    return &MetaRepository{db: db}
}

func (r *MetaRepository) Get(key string) (model.Meta, error) {
    var m model.Meta
    if err := r.db.First(&m, "key = ?", key).Error; err != nil {
        return model.Meta{}, err
    }
    return m, nil
}

func (r *MetaRepository) Set(key, value string) error {
    m := model.Meta{Key: key, Value: value}
    return r.db.Save(&m).Error
}


