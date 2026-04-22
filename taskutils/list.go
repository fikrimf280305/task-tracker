package taskutils

import (
	"fmt"
	"time"
	"strconv"
)

func listAll(store *TaskStore) {
	found := false

	for _, task := range store.Tasks {
		fmt.Print("Id: " + strconv.Itoa(task.Id) + "\nDescription: " + task.Description + "\nStatus: " + string(task.Status) + "\nCreated at: " + task.CreatedAt.Format(time.RFC3339) + "\nUpdated at: " + task.UpdatedAt.Format(time.RFC3339) + "\n\n")

		found = true
	}

	if !found {
		fmt.Println("None")
	}
}

func listWithOption(store *TaskStore, option string) {
	found := false

	status, err := ParseStatus(option)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, task := range store.Tasks {
		if task.Status == status {
			fmt.Print("Id: " + strconv.Itoa(task.Id) + "\nDescription: " + task.Description + "\nStatus: " + string(task.Status) + "\nCreated at: " + task.CreatedAt.Format(time.RFC3339) + "\nUpdated at: " + task.UpdatedAt.Format(time.RFC3339) + "\n\n")

			found = true
		}
	}

	if !found {
		fmt.Println("None")
	}
}

func List(option string) error {
	return withLock(func() error {
		store, err := loadStore()
		if err != nil {
			return err
		}

		if option == "" {
			listAll(store)
		} else {
			listWithOption(store, option)
		}

		return nil
	})
}