package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/govnocods/TaskManager/models"
)

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var id int = 0
	var req models.Task

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}
	models.Mu.Lock()
	task := models.Task{Id: id, Text: req.Text, Completed: false}
	models.Tasks = append(models.Tasks, task)
	id++
	models.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Tasks)
}
