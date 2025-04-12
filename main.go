package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

const filename = "todo.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <add|list|done|delete> [args]")
		return
	}

	cmd := os.Args[1]
	tasks := loadTasks() //

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task description.")
			return
		}
		desc := os.Args[2]
		addTask(&tasks, desc)
	case "list":
		listTasks(tasks)
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Provide task ID to mark done.")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		markDone(&tasks, id)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Provide task ID to delete.")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		deleteTask(&tasks, id)
	default:
		fmt.Println("Unknown command:", cmd)
	}
	saveTasks(tasks)
}

func loadTasks() []Task {
	var tasks []Task
	data, err := os.ReadFile(filename)
	if err != nil {
		return tasks
	}
	json.Unmarshal(data, &tasks)
	return tasks
}
func addTask(tasks *[]Task, desc string) {
	newID := len(*tasks) + 1
	task := Task{ID: newID, Description: desc, Done: false}
	*tasks = append(*tasks, task)
	fmt.Println("Added:", task.Description)
}

func listTasks(tasks []Task) {
	for _, task := range tasks {
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Description)
	}
}

func markDone(tasks *[]Task, id int) {
	for i := range *tasks {
		if (*tasks)[i].ID == id {
			(*tasks)[i].Done = true
			fmt.Println("Marked task done:", (*tasks)[i].Description)
			return
		}
	}
	fmt.Println("Task not found.")
}
func deleteTask(tasks *[]Task, id int) {
	for i, task := range *tasks {
		if task.ID == id {
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			fmt.Println("Deleted task:", task.Description)
			return
		}
	}
	fmt.Println("Task not found.")
}

func saveTasks(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", " ")
	os.WriteFile(filename, data, 0644) // 0 means octal notation , 6 means User ( Read and Write),first 4 means Group (Read),second 4 means other (Read)

}
