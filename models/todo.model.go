package models

type Todo struct {
	Id          int    `json:"id" form:"id"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Completed   bool   `json:"completed" form:"completed"`
}
