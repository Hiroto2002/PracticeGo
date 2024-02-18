package controllers

import (
    "strconv"
    "practiceGo/app/src/interfaces/database"
    "practiceGo/app/src/usecase"
)

type TodosController struct {
    Interactor usecase.TodoInteractor
}

func NewTodosController(db database.DB) *TodosController {
    return &TodosController{
        Interactor: usecase.TodoInteractor{
            DB: &database.DBRepository{ DB: db },
            Todo: &database.TodoRepository{},
        },
    }
}
// controller *TodosControllerはメソッドのレシーバで構造体（または型）にメソッドを「紐付け」ることができます。
func (controller *TodosController) Get(c Context) {

    // strconv.Atoi: 文字列を数値に変換する
    id, _ := strconv.Atoi(c.Param("id"))

    todo, res := controller.Interactor.Get(id)
    if res.Error != nil {
        c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
        return
    }
    c.JSON(res.StatusCode, NewH("success", todo))
}