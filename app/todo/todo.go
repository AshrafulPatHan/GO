package main
import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Task string `json:"task"`
	Done bool   `json:"done"`
}

var fileName = "tasks.json"

// load task from file
func loadTasks() []Todo {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return []Todo{}
	}

	var tasks []Todo
	json.Unmarshal(file, &tasks)
	return tasks
}

// save task to file
func saveTasks(tasks []Todo) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(fileName, data, 0644)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : todo [add/list/done]")
		return
	}

	command := os.Args[1]
	tasks := loadTasks()

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task name!")
			return
		}
		task := os.Args[2]
		tasks = append(tasks, Todo{Task: task, Done: false})
		saveTasks(tasks)
		fmt.Println("Task added:", task)

	case "list":
		for i, t := range tasks {
			status := ""
			if t.Done {
				status = "x"
			}
			fmt.Printf("[%s] %d. %s\n", status, i+1, t.Task)
		}

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide task number!")
			return
		}

		index := os.Args[2]

		var id int
		fmt.Sscanf(index, "%d", &id)

		if id <= 0 || id > len(tasks) {
			fmt.Println("Invalid task number!")
			return
		}

		tasks[id-1].Done = true
		saveTasks(tasks)
		fmt.Println("Marked as done:", tasks[id-1].Task)

	default:
		fmt.Println("Unknown command:", command)
	}
}
