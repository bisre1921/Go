package main

import (
	"fmt"
	"os"
)

type Task struct {
	ID        string
	Name      string
	Completed bool
}

var tasks []Task

func showMenu() {
	fmt.Println("Welcome to the To-Do List App!")
	fmt.Println("1. View Tasks")
	fmt.Println("2. Add Task")
	fmt.Println("3. Mark Task as Completed")
	fmt.Println("4. Delete Task")
	fmt.Println("5. Exit")
	fmt.Print("Enter your choice: ")
}

func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range tasks {
		completedStatus := "Not Completed"
		if task.Completed {
			completedStatus = "Completed"
		}
		fmt.Printf("ID: %s | Task: %s | Status: %s\n", task.ID, task.Name, completedStatus)
	}
}

func addTask() {
	var name string
	fmt.Print("Enter task name: ")
	fmt.Scanln(&name)
	taskID := fmt.Sprintf("%d", len(tasks)+1)
	task := Task{ID: taskID, Name: name, Completed: false}
	tasks = append(tasks, task)
	fmt.Println("Task added successfully.")
}

func markTaskAsCompleted() {
	var taskID string
	fmt.Print("Enter task ID to mark as completed: ")
	fmt.Scanln(&taskID)

	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Completed = true
			fmt.Println("Task marked as completed.")
			return
		}
	}
	fmt.Println("Task not found.")
}

func deleteTask() {
	var taskID string
	fmt.Print("Enter task ID to delete: ")
	fmt.Scanln(&taskID)

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted successfully.")
			return
		}
	}
	fmt.Println("Task not found.")
}

func main() {
	for {
		showMenu()

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			viewTasks()
		case 2:
			addTask()
		case 3:
			markTaskAsCompleted()
		case 4:
			deleteTask()
		case 5:
			fmt.Println("Exiting the app. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
