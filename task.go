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
	CreatedAt   string
	CompletedAt string
}

func (c *config) Add(task string) {
	item := Item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now().Format("Mon 2 Jan 2006 15:04"),
		CompletedAt: time.Time{}.Format("Mon 2 Jan 2006 15:04"),
	}

	c.Tasks = append(c.Tasks, item)
}

func (c *config) Store(filename string) error {
	data, err := json.MarshalIndent(c.Tasks, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (c *config) Load(filename string) error {
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

	if err = json.Unmarshal(file, &c.Tasks); err != nil {
		return err
	}

	return nil
}

func (c *config) CountPending() int {
	total := 0

	for _, item := range c.Tasks {
		if !item.Done {
			total++
		}
	}

	return total
}
