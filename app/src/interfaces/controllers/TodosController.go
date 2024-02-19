package controllers

import (
    // "fmt"
    "net/http"
    "practiceGo/app/src/interfaces/database"
    "practiceGo/app/src/interfaces/dto"
    "practiceGo/app/src/usecase"
    "strconv"

    "github.com/gin-gonic/gin"
)

type TodosController struct {
    Interactor usecase.TodoInteractor
}

func NewTodosController(db database.DB) *TodosController {
    return &TodosController{
        Interactor: usecase.TodoInteractor{
            DB:   &database.DBRepository{DB: db},
            Todo: &database.TodoRepository{},
        },
    }
}

func (controller *TodosController) Get(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    todo, res := controller.Interactor.Get(id)
    if res.Error != nil {
        c.JSON(res.StatusCode, gin.H{"error": res.Error.Error()})
        return
    }
    c.JSON(res.StatusCode, gin.H{"message": "success", "data": todo})
}

func (controller *TodosController) GetAll(c *gin.Context) {
    todos, res := controller.Interactor.GetAll()
    if res.Error != nil {
        c.JSON(res.StatusCode, gin.H{"error": res.Error.Error()})
        return
    }
    c.JSON(res.StatusCode, gin.H{"message": "success", "data": todos})
}





func (controller *TodosController) Create(c *gin.Context) {
    var todoDto dto.TodoForCreate

    if err := c.BindJSON(&todoDto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    createdTodo, res := controller.Interactor.Create(todoDto)
    if res.Error != nil {
        c.JSON(res.StatusCode, gin.H{"error": res.Error.Error()})
        return
    }
    c.JSON(res.StatusCode, gin.H{"message": "success", "data": createdTodo})
}

func (controller *TodosController) Update(c *gin.Context) {
    var todoDto dto.TodoForUpdate

    if err := c.BindJSON(&todoDto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    updatedTodo, res := controller.Interactor.Update(todoDto)
    if res.Error != nil {
        c.JSON(res.StatusCode, gin.H{"error": res.Error.Error()})
        return
    }
    c.JSON(res.StatusCode, gin.H{"message": "success", "data": updatedTodo})
}

func (controller *TodosController) Delete(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    res := controller.Interactor.Delete(id)
    if res.Error != nil {
        c.JSON(res.StatusCode, gin.H{"error": res.Error.Error()})
        return
    }
    c.JSON(res.StatusCode, gin.H{"message": "success"})
}
