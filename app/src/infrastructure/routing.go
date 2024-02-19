package infrastructure

import (
	"github.com/gin-gonic/gin"
    "practiceGo/app/src/interfaces/controllers"
)

type Routing struct {
    DB *DB
    Http *Http
    Gin *gin.Engine
    Port string
}

func NewRouting(db *DB,http *Http) *Routing {
    c := NewConfig() 
    r := &Routing{
        DB: db,
        Http: http,
        Gin: gin.Default(),
        Port: c.Routing.Port,
    }
    r.setRouting()
    return r
}

func (r *Routing) setRouting() {
    todosController := controllers.NewTodosController(r.DB)
    r.Gin.GET("/todos", func (c *gin.Context) { todosController.GetAll(c) })
    r.Gin.GET("/todos/:id", func (c *gin.Context) { todosController.Get(c) })
    r.Gin.POST("/todos", func (c *gin.Context) { todosController.Create(c) })
    r.Gin.PUT("/todos/:id", func (c *gin.Context) { todosController.Update(c) })
    r.Gin.DELETE("/todos/:id", func (c *gin.Context) { todosController.Delete(c) })
}

func (r *Routing) Run() {
    r.Gin.Run(r.Port)
}