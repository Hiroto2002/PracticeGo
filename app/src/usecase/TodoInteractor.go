// clean architecture における usecase の実装,service 層
package usecase

import (
    "practiceGo/app/src/interfaces/dto"
)

type TodoInteractor struct {
    DB DBRepository
    Todo TodoRepository
}

func (interactor *TodoInteractor) Get(id int) (todoDto dto.TodoForGet, resultStatus *ResultStatus) {
    db := interactor.DB.Connect()
    // Todos の取得
    foundTodo, err := interactor.Todo.FindByID(db, id)
    if err != nil {
        return dto.TodoForGet{}, NewResultStatus(404, err)
    }
    todoDto = dto.TodoForGet{
        ID:          int(foundTodo.ID),
        Title:       foundTodo.Title,
        Description: foundTodo.Description,
        IsCompleted: foundTodo.IsCompleted,
    }
    return todoDto, NewResultStatus(200, nil)
}