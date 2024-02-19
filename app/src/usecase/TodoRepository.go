package usecase

import (
    "practiceGo/app/src/domain"
    "gorm.io/gorm"
)

type TodoRepository interface {
    FindByID(db *gorm.DB, id int) (todo domain.Todos, err error)
    FindAll(db *gorm.DB) (todos []domain.Todos, err error)
    Create(db *gorm.DB, todo *domain.Todos) (*domain.Todos, error)
    Update(db *gorm.DB, todo *domain.Todos) (*domain.Todos, error) 
    Delete(db *gorm.DB, id int) (err error)
}