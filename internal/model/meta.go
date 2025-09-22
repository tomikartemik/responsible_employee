package model

type Meta struct {
    Key   string `gorm:"primaryKey;unique" json:"key"`
    Value string `json:"value"`
}


