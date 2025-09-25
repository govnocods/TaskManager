package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/govnocods/TaskManager/models"
	"github.com/govnocods/TaskManager/utils/id"
)

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

	task := models.Task{
		Id:        id.NextID(),
		Text:      req.Text,
		Completed: false,
	}

	models.Tasks.Store(task.Id, task)

	var allTasks []models.Task
	models.Tasks.Range(func(key, value any) bool {
		if task, ok := value.(models.Task); ok {
			allTasks = append(allTasks, task)
		}
		return true
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allTasks)
}

func DelTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}

	var req models.Task

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON parsing error", http.StatusBadRequest)
		return
	}
	models.Tasks.Delete(req.Id)
}

func EditTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	var req models.Task
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	if _, ok := models.Tasks.Load(req.Id); ok {
		models.Tasks.Store(req.Id, req)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(req)
		return
	}
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tasks []models.Task
	models.Tasks.Range(func(key, value any) bool {
		if task, ok := value.(models.Task); ok {
			tasks = append(tasks, task)
		}
		return true
	})
	json.NewEncoder(w).Encode(tasks)
}
