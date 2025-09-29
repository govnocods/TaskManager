package db

import (
	"log"

	"github.com/govnocods/TaskManager/models"
)

func AddTask(task models.Task) {
	query := `INSERT INTO tasks (task, completed) VALUES(?, ?)`

	_, err := Connect().Exec(query, task.Text, task.Completed)
	if err != nil {
		log.Fatal(err)
	}
}

func EditTask(task models.Task) {
	query := `UPDATE tasks SET task = ?, completed = ? WHERE id = ?`

	_, err := Connect().Exec(query, task.Text, task.Completed, task.Id)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteTask(task models.Task) {

	query := `DELETE FROM tasks WHERE id = ?`

	_, err := Connect().Exec(query, task.Id)
	if err != nil {
		log.Fatal(err)
	}
}

func GetTasks() []models.Task {
	query := `SELECT * FROM tasks`

	var tasks []models.Task

	rows, err := Connect().Query(query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.Id, &task.Text, &task.Completed); err != nil {
			log.Printf("Error scanning task row: %v", err)
			continue
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		return nil
	}

	return tasks
}
