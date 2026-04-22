package taskutils

import (
	"errors"
	"os"
	"time"
)

func Mark(id int, status Status) error {
	return withLock(func() error {
		if id == 0 {
			return errors.New("Task Id required")
		}

		if status == "" {
			return errors.New("Task new status required")
		}

		store, err := loadStore()
		if err != nil {
			return err
		}

		found := false

		for i, task := range store.Tasks {
			if task.Id == id {
				store.Tasks[i].Status = status
				store.Tasks[i].UpdatedAt = time.Now()
				found = true
				break
			}
		}

		if !found {
			return os.ErrNotExist
		}

		return saveStore(store)
	})
}