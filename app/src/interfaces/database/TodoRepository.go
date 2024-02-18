package database

import (
    "errors"
    "gorm.io/gorm"
    "practiceGo/app/src/domain"
)

type TodoRepository struct {}

func (repo *TodoRepository) FindByID(db *gorm.DB, id int) (todo domain.Todos, err error) {
    todo = domain.Todos{}
    db.First(&todo, id)
    if todo.ID <= 0 {
        return domain.Todos{}, errors.New("todo is not found")
    }

    return todo, nil
}