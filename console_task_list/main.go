package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		status, tasks, err := readTasks()
		if err != nil && !os.IsNotExist(err) {
			fmt.Println("Error reading tasks:", err)
			return
		}

		if !status {
			tasks = []Task{}
		}

		displayMainMenu()

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			displayTasks(tasks)
		case "2":
			tasks, err = handleAddTask(tasks)
			if err != nil {
				fmt.Println("Error adding task:", err)
			}
		case "3":
			tasks, err = handleUpdateTask(tasks)
			if err != nil {
				fmt.Println("Error updating task:", err)
			}
		case "4":
			tasks, err = handleDeleteTask(tasks)
			if err != nil {
				fmt.Println("Error deleting task:", err)
			}
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option selected. Exiting...")
			return
		}
	}
}

func displayMainMenu() {
	fmt.Println("\nMain Menu")
	fmt.Println("---------")
	fmt.Println("1. View Tasks")
	fmt.Println("2. Add Task")
	fmt.Println("3. Update Task")
	fmt.Println("4. Delete Task")
	fmt.Println("5. Exit")
	fmt.Println("---------")
	fmt.Print("Enter an option: ")
}

// displayTasks prints all tasks in a readable format
func displayTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to display.")
		return
	}
	for _, task := range tasks {
		fmt.Printf("Task ID: %d\nPriority: %s\nDue Date: %s\nDescription: %s\nStatus: %s\n\n",
			task.ID, task.Priority, task.DueDate, task.Description, task.Status)
	}
}

// handleAddTask prompts the user for task details, adds a new task, and returns the updated slice
func handleAddTask(tasks []Task) ([]Task, error) {
	var (
		description, priorityOption, dueDate, status string
	)

	fmt.Print("Enter the description of the task: ")
	fmt.Scanln(&description)
	fmt.Println("Select the priority of the task:")
	fmt.Println("1. High")
	fmt.Println("2. Medium")
	fmt.Println("3. Low")
	fmt.Print("Enter an option: ")
	fmt.Scanln(&priorityOption)

	priority := map[string]string{"1": High, "2": Medium, "3": Low}[priorityOption]

	fmt.Print("Enter the due date of the task (YYYY-MM-DD): ")
	fmt.Scanln(&dueDate)

	status = Pending // Default status

	task := createTask(priority, dueDate, description, status)
	return addTask(tasks, task)
}

// handleUpdateTask prompts the user to update specific fields of a task
func handleUpdateTask(tasks []Task) ([]Task, error) {
	var (
		id    int
		field string
	)

	fmt.Print("Enter the ID of the task you would like to update: ")
	fmt.Scanln(&id)
	task := getTaskByID(tasks, id)
	if task.ID == 0 {
		fmt.Println("Task not found.")
		return tasks, nil // No task found, no update performed
	}

	fmt.Println("What would you like to update?")
	fmt.Println("1. Priority")
	fmt.Println("2. Due Date")
	fmt.Println("3. Description")
	fmt.Println("4. Status")
	fmt.Println("5. Back to Main Menu")
	fmt.Print("Enter an option: ")
	fmt.Scanln(&field)

	switch field {
	case "1":
		fmt.Print("Enter the new priority (High, Medium, Low): ")
		var priority string
		fmt.Scanln(&priority)
		return updateTask(tasks, task.ID, priority, task.DueDate, task.Description, task.Status)
	case "2":
		fmt.Print("Enter the new due date (YYYY-MM-DD): ")
		var dueDate string
		fmt.Scanln(&dueDate)
		return updateTask(tasks, task.ID, task.Priority, dueDate, task.Description, task.Status)
	case "3":
		fmt.Print("Enter the new description: ")
		var description string
		fmt.Scanln(&description)
		return updateTask(tasks, task.ID, task.Priority, task.DueDate, description, task.Status)
	case "4":
		fmt.Print("Enter the new status (Pending, In Progress, Completed): ")
		var status string
		fmt.Scanln(&status)
		return updateTask(tasks, task.ID, task.Priority, task.DueDate, task.Description, status)
	case "5":
		// Back to Main Menu, no changes
	default:
		fmt.Println("Invalid option selected.")
	}
	return tasks, nil
}

// handleDeleteTask prompts the user for an ID and deletes the corresponding task
func handleDeleteTask(tasks []Task) ([]Task, error) {
	var id int
	fmt.Print("Enter the ID of the task you would like to delete: ")
	fmt.Scanln(&id)
	return deleteTask(tasks, id)
}

// getTaskByID returns a task by its ID
func getTaskByID(tasks []Task, id int) Task {
	for _, task := range tasks {
		if task.ID == id {
			return task
		}
	}
	return Task{} // Return an empty Task if not found
}
