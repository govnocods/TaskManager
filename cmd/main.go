package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/govnocods/TaskManager/internal/db"
	"github.com/govnocods/TaskManager/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	db.Connect()
	defer db.CloseDB()

	mux.HandleFunc("/get", handlers.GetTaskHandler)
	mux.HandleFunc("/add", handlers.AddTaskHandler)
	mux.HandleFunc("/edit", handlers.EditTaskHandler)
	mux.HandleFunc("/del", handlers.DelTaskHandler)

	fs := http.FileServer(http.Dir("./web"))

	mux.Handle("/", fs)
	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
