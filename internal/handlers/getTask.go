package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/govnocods/TaskManager/models"
)

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Tasks)
}
