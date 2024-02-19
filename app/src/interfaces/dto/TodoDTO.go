package dto
import(
    "practiceGo/app/src/domain"
)

type TodoForGet struct {
    ID int `json:"id"`
    Title string `json:"title"`
    Description string `json:"description"`
    IsCompleted bool `json:"isCompleted"`
}

type TodoForCreate struct {
    Title string `json:"title"`
    Description string `json:"description"`
    IsCompleted bool `json:"isCompleted"`
}

type TodoForUpdate struct {
    ID int `json:"id"`
    Title string `json:"title"`
    Description string `json:"description"`
    IsCompleted bool `json:"isCompleted"`
}

func (todoDto *TodoForCreate) ToEntity() (todo domain.Todos) {
    return domain.Todos{
        Title:       todoDto.Title,
        Description: todoDto.Description,
        IsCompleted: todoDto.IsCompleted,
    }
}

func (todoDto *TodoForUpdate) ToEntity() (todo domain.Todos) {
    return domain.Todos{
        Title:       todoDto.Title,
        Description: todoDto.Description,
        IsCompleted: todoDto.IsCompleted,
    }
}