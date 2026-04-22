package taskutils

import (
	"time"
	"encoding/json"
	"os"
	"github.com/gofrs/flock"
	"fmt"
)

type Task struct {
	Id 			int 		`json:"id"`
	Description string 		`json:"description"`
	Status 		Status 		`json:"status"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}

type Status string

const (
	Unknown		Status = "unknown"
	Todo 		Status = "todo"
	InProgress 	Status = "in-progress"
	Done 		Status = "done"
)

type TaskStore struct {
	LastID 	int 	`json:"last_id"`
	Tasks 	[]Task 	`json:"tasks"`
}

const fileName = "tasks.json"

const lockFile = "tasks.lock"

func withLock(fn func() error) error {
	lock := flock.New(lockFile)

	if err := lock.Lock(); err != nil {
		return err
	}
	
	defer lock.Unlock()

	return fn()
}

func loadStore() (*TaskStore, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return &TaskStore{LastID: 0, Tasks: []Task{}}, nil
		}
		return nil, err
	}

	var store TaskStore
	if err := json.Unmarshal(data, &store); err != nil {
		return nil, err
	}

	return &store, nil
}

func saveStore(store *TaskStore) error {
	tmpFile := fileName + ".tmp"

	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(tmpFile, data, 0644); err != nil {
		return err
	}

	return os.Rename(tmpFile, fileName)
}

func ParseStatus(s string) (Status, error) {
	switch s {
	case "todo":
		return Todo, nil
	case "in-progress":
		return InProgress, nil
	case "done":
		return Done, nil
	default:
		return Unknown, fmt.Errorf("invalid status")
	}
}
