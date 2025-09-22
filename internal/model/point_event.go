package model

import "time"

type PointEvent struct {
    ID        int        `gorm:"primaryKey;uniqueIndex;autoIncrement" json:"id"`
    UserID    string     `gorm:"not null" json:"userId"`
    Source    string     `gorm:"not null" json:"source"`
    Points    int        `gorm:"not null" json:"points"`
    CreatedAt time.Time  `gorm:"not null;autoCreateTime" json:"createdAt"`
    TaskID    *string    `json:"taskId,omitempty"`
    ReportID  *string    `json:"reportId,omitempty"`
}

type UserPointsBreakdown struct {
    UserID              string         `json:"userId"`
    Login               string         `json:"login"`
    FullName            string         `json:"fullName"`
    MonthlyTotal        int            `json:"monthlyTotal"`
    YearlyTotal         int            `json:"yearlyTotal"`
    MonthlyBySource     map[string]int `json:"monthlyBySource"`
    YearlyBySource      map[string]int `json:"yearlyBySource"`
}


