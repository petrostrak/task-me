package main

import (
	"fmt"
	"time"
)

func (c *config) addButton() func() {
	return func() {
		c.Add(c.TaskEntry.Text)
		c.Store(TASKS_FILE)

		c.TaskEntry.Text = ""
		c.TaskEntry.Refresh()
		c.refreshPendings()
	}
}

func (c *config) completeButton() func() {
	return func() {
		var TempData []Item

		for _, i := range c.Tasks {
			if c.TaskLabel.Text == i.Task {

				item := Item{
					Task:        i.Task,
					Done:        true,
					CreatedAt:   i.CreatedAt,
					CompletedAt: time.Now(),
				}

				TempData = append(TempData, item)
			} else {
				TempData = append(TempData, i)
			}
		}

		c.Tasks = TempData
		c.Store(TASKS_FILE)

		c.TaskEntry.Text = ""
		c.TaskEntry.Refresh()
		c.refreshPendings()
	}
}

func (c *config) deleteButton() func() {
	return func() {
		var TempData []Item

		for _, i := range c.Tasks {
			if c.TaskLabel.Text != i.Task {
				TempData = append(TempData, i)
			}
		}

		c.Tasks = TempData
		c.Store(TASKS_FILE)
		c.refreshPendings()
	}
}

func (c *config) refreshPendings() {
	c.Counter = c.CountPending()
	c.Pendings.Set(fmt.Sprintf("You have %d pending task(s)", c.Counter))
}
