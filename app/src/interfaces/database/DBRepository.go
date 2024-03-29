package database

import (
    "gorm.io/gorm"
)

type DBRepository struct {
    DB DB
}

func (db *DBRepository) Begin() *gorm.DB {
    return db.DB.Begin()
}

func (db *DBRepository) Connect() *gorm.DB {
    return db.DB.Connect()
}
