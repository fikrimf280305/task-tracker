package main

import (
	"flag"
	"fmt"
	"os"
	"task-tracker/taskutils"
)

func main() {
	// Add Command
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addIdTask := addCmd.String("name", "", "Task Name")

	// Update Command
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateTaskId := updateCmd.Int("id", 0, "Task ID")
	updateTaskName := updateCmd.String("name", "", "Task Name")

	// Delete Command
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteTaskId := deleteCmd.Int("id", 0, "Task ID")

	// Mark Command
	markCmd := flag.NewFlagSet("mark", flag.ExitOnError)
	markId := markCmd.Int("id", 0, "Task Id")
	markOption := markCmd.String("option", "", "Mark Option: \"in-progress\" or \"done\"")

	// List Command
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listOption := listCmd.String("option", "", "List Option: \"todo\", \"in-progress\", or \"done\"")

	if len(os.Args) < 2 {
		fmt.Println("expected commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])

		if *addIdTask == "" {
			fmt.Println("Task name is required")
			os.Exit(1)
		}

		err := taskutils.Add(*addIdTask)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("Task added successfully")
	case "update":
		updateCmd.Parse(os.Args[2:])

		if *updateTaskId == 0 {
			fmt.Println("Task Id is required")
			os.Exit(1)
		}

		if *updateTaskName == "" {
			fmt.Println("New task name is required")
			os.Exit(1)
		}

		err := taskutils.Update(*updateTaskId, *updateTaskName)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("Task updated successfully")
	case "delete":
		deleteCmd.Parse(os.Args[2:])

		if *deleteTaskId == 0 {
			fmt.Println("Task id required")
			os.Exit(1)
		}

		err := taskutils.Delete(*deleteTaskId)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("Task deleted successfully")
	case "mark":
		markCmd.Parse(os.Args[2:])

		if *markId == 0 {
			fmt.Println("task id required")
			os.Exit(1)
		}

		if *markOption == "" {
			fmt.Println("status required")
			os.Exit(1)
		}

		status, err := taskutils.ParseStatus(*markOption)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		err = taskutils.Mark(*markId, status)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("Task not found")
			} else {
				fmt.Println("Error:", err)
			}

			os.Exit(1)
		}

		fmt.Println("Task Updated Successfully")
	case "list":
		listCmd.Parse(os.Args[2:])

		err := taskutils.List(*listOption)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
