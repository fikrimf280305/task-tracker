package taskutils

import (
	"errors"
	"time"
)

func Add(name string) error {
	return withLock(func() error {
		if name == "" {
			return errors.New("task name cannot be empty")
		}

		store, err := loadStore()
		if err != nil {
			return err
		}

		store.LastID++
		now := time.Now()

		newTask := Task{
			Id:          store.LastID,
			Description: name,
			Status:      Todo,
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		store.Tasks = append(store.Tasks, newTask)

		return saveStore(store)
	})
}
