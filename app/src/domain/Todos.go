package domain

import "gorm.io/gorm"

type Todos struct {
    gorm.Model
    Title        string `gorm:"column:title"`
    Description  string `gorm:"column:description"`
    IsCompleted  bool   `gorm:"column:is_completed"`
}
