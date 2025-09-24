package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/govnocods/TaskManager/models"
)

func DelTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}

	var req models.Task

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON parsing error", http.StatusBadRequest)
		return
	}

	models.Mu.Lock()
	for i, t := range models.Tasks {
		if t.Id == req.Id {
			models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
			break
		}
	}
	models.Mu.Unlock()
}
