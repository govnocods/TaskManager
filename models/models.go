package models

type Task struct {
	Id        int64  `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}
