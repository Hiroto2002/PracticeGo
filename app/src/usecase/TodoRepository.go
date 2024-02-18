package usecase

import (
    "practiceGo/app/src/domain"
    "gorm.io/gorm"
)

type TodoRepository interface {
    FindByID(db *gorm.DB, id int) (todo domain.Todos, err error)
}