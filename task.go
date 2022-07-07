package main

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Tasks []Item

func (t *Tasks) Add(task string) {
	item := Item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, item)
}

func (t *Tasks) Store(filename string) error {
	data, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (t *Tasks) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return errors.New("file is empty")
	}

	if err = json.Unmarshal(file, t); err != nil {
		return err
	}

	return nil
}

func (t *Tasks) CountPending() int {
	total := 0

	for _, item := range *t {
		if !item.Done {
			total++
		}
	}

	return total
}
