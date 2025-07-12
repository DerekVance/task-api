package models

type ErrorResponse struct {
    Error string `json:"error"`
}

type Task struct {
	ID        int    `json:"id" example:"1"`
	Title     string `json:"title" example:"Learn Piano"`
	Completed bool   `json:"completed" example:"false"`
}