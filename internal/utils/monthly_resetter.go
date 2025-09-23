package utils

import (
    "fmt"
    "time"
    "gorm.io/gorm"
    "responsible_employee/internal/model"
)

// MonthlyResetter resets monthly points for all users on the 1st day of each month.
// It uses the Meta table to guard against multiple resets within the same month (key: last_monthly_reset = YYYY-MM).
type MonthlyResetter struct {
    db     *gorm.DB
    quit   chan struct{}
}

func NewMonthlyResetter(db *gorm.DB) *MonthlyResetter {
    return &MonthlyResetter{db: db, quit: make(chan struct{})}
}

func (r *MonthlyResetter) Start() {
    go func() {
        ticker := time.NewTicker(24 * time.Hour)
        defer ticker.Stop()
        for {
            select {
            case <-ticker.C:
                r.tryReset()
            case <-r.quit:
                return
            }
        }
    }()
}

func (r *MonthlyResetter) Stop() { close(r.quit) }

func (r *MonthlyResetter) tryReset() {
    now := time.Now()
    if now.Day() != 1 {
        return
    }

    monthKey := now.Format("2006-01")
    var meta model.Meta
    err := r.db.First(&meta, "key = ?", "last_monthly_reset").Error
    if err == nil && meta.Value == monthKey {
        return // already reset this month
    }

    // Сначала сохраняем текущие месячные баллы в last_month_points, затем обнуляем monthly_points
    if err := r.db.Model(&model.User{}).Update("last_month_points", gorm.Expr("monthly_points")).Error; err != nil {
        fmt.Println("monthly reset error (preserve):", err)
        return
    }
    
    if err := r.db.Model(&model.User{}).Update("monthly_points", 0).Error; err != nil {
        fmt.Println("monthly reset error (zero):", err)
        return
    }

    // store meta
    m := model.Meta{Key: "last_monthly_reset", Value: monthKey}
    _ = r.db.Save(&m).Error
}


