package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the TO DO List CLI app!")
	var tasks []string

	for {
		fmt.Println("\nEnter your command (create, read, update, delete): ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "create":
			fmt.Print("Enter task name: ")
			task, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			task = strings.TrimSpace(task)
			if task != "" {
				tasks = append(tasks, task)
				fmt.Println("Task added!")
			} else {
				fmt.Println("Task name cannot be empty!")
			}

		case "read":
			if len(tasks) == 0 {
				fmt.Println("No tasks yest!")
			} else {
				fmt.Println("Your tasks:")
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}

		case "update":
			if len(tasks) == 0 {
				fmt.Println("No tasks to update!")
				continue
			}
			fmt.Print("Enter task number to update: ")
			taskNumStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			taskNumStr = strings.TrimSpace(taskNumStr)

			var taskNum int
			if _, err := fmt.Scanf(taskNumStr, "%d", &taskNum); err != nil || taskNum < 1 || taskNum > len(tasks) {
				fmt.Println("Invalid task number!")
				continue
			}

			removedTask := tasks[taskNum-1]
			tasks = append(tasks[:taskNum-1], tasks[taskNum:]...)
			fmt.Printf("Removed task #%d ('%s') successfully!\n", taskNum, removedTask)

		default:
			fmt.Println("Invalid command! Available command: create, read, update, delete")
		}
	}
}
