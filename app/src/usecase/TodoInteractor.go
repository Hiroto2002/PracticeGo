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
        return dto.TodoForGet{}, NewResultStatus(404, err,)
    }
    todoDto = dto.TodoForGet{
        ID:          int(foundTodo.ID),
        Title:       foundTodo.Title,
        Description: foundTodo.Description,
        IsCompleted: foundTodo.IsCompleted,
    }
    return todoDto, NewResultStatus(200, nil)
}

func (interactor *TodoInteractor) GetAll() (todosDto []dto.TodoForGet, resultStatus *ResultStatus) {
    db := interactor.DB.Connect()
    // Todos の取得
    foundTodos, err := interactor.Todo.FindAll(db)
    if err != nil {
        return []dto.TodoForGet{}, NewResultStatus(404, err)
    }
    for _, todo := range foundTodos {
        todosDto = append(todosDto, dto.TodoForGet{
            ID:          int(todo.ID),
            Title:       todo.Title,
            Description: todo.Description,
            IsCompleted: todo.IsCompleted,
        })
    }
    return todosDto, NewResultStatus(200, nil)
}

func (interactor *TodoInteractor) Create(todoDto dto.TodoForCreate) (*dto.TodoForGet, *ResultStatus) {
    db := interactor.DB.Connect()
    todo := todoDto.ToEntity()

    createdTodo, err := interactor.Todo.Create(db, &todo)
    if err != nil {
        // エラーが発生した場合は、todoオブジェクトにnilを設定
        return nil, NewResultStatus(500, err)
    }
    resultTodo := &dto.TodoForGet{
        ID:          int(createdTodo.ID),
        Title:       createdTodo.Title,
        Description: createdTodo.Description,
        IsCompleted: createdTodo.IsCompleted,
    }
    return resultTodo, NewResultStatus(201, nil)
}

func (interactor *TodoInteractor) Update(todoDto dto.TodoForUpdate) (*dto.TodoForGet, *ResultStatus) {
    db := interactor.DB.Connect()
    todo := todoDto.ToEntity()

    updatedTodo, err := interactor.Todo.Update(db, &todo)
    if err != nil {
        return nil, NewResultStatus(500, err)
    }

    // 更新されたTodoをTodoForGet DTOに変換
    resultTodo := &dto.TodoForGet{
        ID:          int(updatedTodo.ID),
        Title:       updatedTodo.Title,
        Description: updatedTodo.Description,
        IsCompleted: updatedTodo.IsCompleted,
    }

    return resultTodo, NewResultStatus(200, nil)
}

func (interactor *TodoInteractor) Delete(id int) (resultStatus *ResultStatus) {
    db := interactor.DB.Connect()
    // Todos の削除
    err := interactor.Todo.Delete(db, id)
    if err != nil {
        return NewResultStatus(500, err)
    }
    return NewResultStatus(200, nil)
}