package main

import (
	"encoding/json"
	"os"
	"sort"
)

// Priority levels
const (
	High   = "high"
	Medium = "medium"
	Low    = "low"
)

// Task statuses
const (
	InProgress = "in progress"
	Completed  = "completed"
	Pending    = "pending"
)

// Task represents a task with its attributes
type Task struct {
	ID          int    `json:"id"`
	Priority    string `json:"priority"`
	DueDate     string `json:"due_date"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// readTasks checks for savedTasks.json and returns tasks if exists
func readTasks() (bool, []Task, error) {
	if _, err := os.Stat("savedTasks.json"); os.IsNotExist(err) {
		return false, nil, err
	}
	return readTasksFromFile("savedTasks.json")
}

// readTasksFromFile reads tasks from the specified file
func readTasksFromFile(filename string) (bool, []Task, error) {
	file, err := os.Open(filename)
	if err != nil {
		return false, nil, err
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return false, nil, err
	}
	return true, tasks, nil
}

// createTask creates a new Task instance
func createTask(priority, dueDate, description, status string) Task {
	return Task{-1, priority, dueDate, description, status}
}

// sortTasksByID sorts tasks by their ID
func sortTasksByID(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})
}

// addTask adds a new task to the slice and saves it
func addTask(tasks []Task, task Task) ([]Task, error) {
	task.ID = 1
	if len(tasks) > 0 {
		task.ID = tasks[len(tasks)-1].ID + 1
	}
	tasks = append(tasks, task)
	sortTasksByID(tasks)
	return tasks, saveTasksToFile(tasks, "savedTasks.json")
}

// updateTask updates a task by its ID
func updateTask(tasks []Task, id int, priority, dueDate, description, status string) ([]Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i] = Task{id, priority, dueDate, description, status}
			break
		}
	}
	return tasks, saveTasksToFile(tasks, "savedTasks.json")
}

// deleteTask removes a task by its ID
func deleteTask(tasks []Task, id int) ([]Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	sortTasksByID(tasks)
	return tasks, saveTasksToFile(tasks, "savedTasks.json")
}

// saveTasksToFile saves the tasks to the specified file
func saveTasksToFile(tasks []Task, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}
