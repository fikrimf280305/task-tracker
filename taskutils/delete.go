package taskutils

import (
	"errors"
	"os"
)

func Delete(id int) error {
	return withLock(func() error {
		if id == 0 {
			return errors.New("Missing task id")
		}

		store, err := loadStore()
		if err != nil {
			return err
		}

		found := false
		newTasks := make([]Task, 0, len(store.Tasks))

		for _, task := range store.Tasks {
			if task.Id == id {
				found = true
				continue
			}

			newTasks = append(newTasks, task)
		}

		if !found {
			return os.ErrNotExist
		}

		store.Tasks = newTasks

		return saveStore(store)
	})
}