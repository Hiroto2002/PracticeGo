package dto

type TodoForGet struct {
    ID int `json:"id"`
    Title string `json:"title"`
    Description string `json:"description"`
    IsCompleted bool `json:"isCompleted"`
}