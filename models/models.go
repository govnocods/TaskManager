package models

import (
	"sync"
)

type Task struct {
	Id        int64  `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

var (
	Tasks sync.Map
)
