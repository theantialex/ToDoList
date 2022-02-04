package models

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type MarkRequest struct {
	Done bool `json:"done"`
}
