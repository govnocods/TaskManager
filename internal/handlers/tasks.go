package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/govnocods/TaskManager/internal/db"
	"github.com/govnocods/TaskManager/models"
)

func respondJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.Task

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	db.AddTask(req)
	tasks := db.GetTasks()

	respondJSON(w, tasks)
}

func DelTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var req models.Task

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON parsing error", http.StatusBadRequest)
		return
	}

	db.DeleteTask(req)
	respondJSON(w, req)
}

func EditTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.Task
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}
	db.EditTask(req)
	w.WriteHeader(http.StatusOK)
	respondJSON(w, req)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	tasks := db.GetTasks()
	respondJSON(w, tasks)
}
