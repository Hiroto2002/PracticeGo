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

func (repo *TodoRepository) FindAll(db *gorm.DB) (todos []domain.Todos, err error) {
    todos = []domain.Todos{}
    db.Find(&todos)
    if len(todos) <= 0 {
        return []domain.Todos{}, errors.New("todos are not found")
    }

    return todos, nil
}

func (repo *TodoRepository) Create(db *gorm.DB, todo *domain.Todos)(*domain.Todos, error)  {
    // データベースでの登録処理
    if err := db.Create(todo).Error; err != nil {
        return nil, err
    }
    return todo, nil // 登録されたtodoを返す
}

func (r *TodoRepository) Update(db *gorm.DB, todo *domain.Todos) (*domain.Todos, error) {
    // データベースでの更新処理
    if err := db.Save(todo).Error; err != nil {
        return nil, err
    }
    return todo, nil // 更新されたtodoを返す
}

func (repo *TodoRepository) Delete(db *gorm.DB, id int) error {
    result := db.Delete(&domain.Todos{}, "id = ?", id) // 削除条件を指定
    return result.Error
}

