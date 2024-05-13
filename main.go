package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Text      string
	Completed bool
}

const (
	dataPath = "data"
)

func main() {
	tasks := []Task{}

	for {
		createFolderToSave()
		showMenu()

		switch option := getUserInput("Enter your choice: "); option {
		case "1":
			showTasks(tasks)
		case "2":
			addTask(&tasks)
		case "3":
			markTaskCompleted(&tasks)
		case "4":
			saveTasksToFile(tasks)
		case "5":
			fmt.Println("Exit todo app")
			return
		default:
			fmt.Println("Invalid choice. Please try again!")
		}
	}
}

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func createFolderToSave() {
	_, err := os.Stat(dataPath)
	if err == nil {
		return
	}
	if err := os.Mkdir(dataPath, os.ModePerm); err != nil {
		panic(err)
	}
}

func showMenu() {
	fmt.Println("--------------------------------")
	fmt.Println("| Menu:                        |")
	fmt.Println("| 1. Show Tasks                |")
	fmt.Println("| 2. Add Task                  |")
	fmt.Println("| 3. Mark Task as Completed    |")
	fmt.Println("| 4. Save Tasks to File        |")
	fmt.Println("| 5. Exit                      |")
	fmt.Println("--------------------------------")
}

func showTasks(tasks []Task) {
	fmt.Println("--- Show tasks ---")
	if len(tasks) == 0 {
		fmt.Println("No tasks available")
		return
	}
	fmt.Println("Tasks: ")
	for i, task := range tasks {
		status := " "
		if task.Completed {
			status = "x"
		}
		fmt.Printf("%d. [%s] %s \n", i+1, status, task.Text)
	}
}

func addTask(tasks *[]Task) {
	taskText := getUserInput("Enter task description: ")
	*tasks = append(*tasks, Task{Text: taskText})
	fmt.Println("Task added")
}

func markTaskCompleted(tasks *[]Task) {
	showTasks(*tasks)
	taskIndexStr := getUserInput("Enter task index to mark as completed: ")

	taskIndex, err := strconv.Atoi(taskIndexStr)
	if err != nil || taskIndex < 1 || taskIndex > len(*tasks) {
		fmt.Println("Invalid task number. Please try again")
		return
	}

	(*tasks)[taskIndex-1].Completed = true
	fmt.Println("Task marked as completed")
}

func saveTasksToFile(tasks []Task) {
	fileName := getUserInput("Enter filename to save: ")
	filePath := fmt.Sprintf("%s/%s.txt", dataPath, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	defer file.Close()

	for i, task := range tasks {
		status := " "
		if task.Completed {
			status = "x"
		}
		file.WriteString(fmt.Sprintf("%d. [%s] %s \n", i+1, status, task.Text))
	}
	fmt.Printf("Tasks saved to file '%s.txt'\n", fileName)
}
