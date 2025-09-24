package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/govnocods/TaskManager/models"
)

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

	models.Mu.Lock()
	for i, t := range models.Tasks {
		if t.Id == req.Id {
			models.Tasks[i].Text = req.Text
		}
	}
	models.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Tasks)
}
