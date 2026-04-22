package taskutils

import (
	"errors"
	"os"
	"time"
)

func Update(id int, name string) error {
	return withLock(func() error {
		if id == 0 {
			return errors.New("Missing task id")
		}

		if name == "" {
			return errors.New("Missing task new name")
		}

		store, err := loadStore()
		if err != nil {
			return err
		}

		found := false

		for i, task := range store.Tasks {
			if task.Id == id {
				store.Tasks[i].Description = name
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
