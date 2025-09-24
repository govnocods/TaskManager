package models

import "sync"

type Task struct {
	Id        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

var (
	Tasks []Task
	Mu    sync.Mutex
)
