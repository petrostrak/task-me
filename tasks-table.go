package main

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2/widget"
	"github.com/petrostrak/task-me/repository"
)

func (c *config) getTasksTable() *widget.Table {
	data := c.getTaskSlice()
}

func (c *config) getTaskSlice() [][]any {
	var slice [][]any

	tasks, err := c.currentTasks()
	if err != nil {
		log.Println(err)
	}

	slice = append(slice, []any{"ID", "Title", "Complete", "Delete"})

	for _, x := range tasks {
		var currentRow []any

		currentRow = append(currentRow, strconv.FormatInt(x.ID, 10))
		currentRow = append(currentRow, x.Title)
		currentRow = append(currentRow, widget.NewButton("Complete", func() {}))
		currentRow = append(currentRow, widget.NewButton("Delete", func() {}))

		slice = append(slice, currentRow)
	}

	return slice
}

func (c *config) currentTasks() ([]repository.Task, error) {
	tasks, err := c.DB.AllTasks()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
