package models

type Note struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	content string `json:"content"`
}
