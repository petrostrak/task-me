package main

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/petrostrak/task-me/repository"
)

func (c *config) addTaskDialog() dialog.Dialog {
	addLableEntry := widget.NewEntry()
	addDescriptionEntry := widget.NewEntry()

	c.AddTasksLableEntry = addLableEntry
	c.AddTasksDescriptionEntry = addDescriptionEntry

	// create a dialog
	addForm := dialog.NewForm(
		"Add Task",
		"Add",
		"Cancel",
		[]*widget.FormItem{
			{Text: "Title of task", Widget: addLableEntry},
			{Text: "Description", Widget: addDescriptionEntry},
		},
		func(valid bool) {
			if valid {
				title := addLableEntry.Text
				description := addDescriptionEntry.Text

				_, err := c.DB.InsertTask(repository.Task{
					Title:       title,
					Description: description,
					Done:        false,
					CreatedAt:   time.Now(),
					CompletedAt: time.Time{},
				})
				if err != nil {
					log.Println(err)
				}

				c.refreshTaskTable()
			}
		}, c.MainWindow)

	// size and show the dialog
	addForm.Resize(fyne.Size{Width: 400})
	addForm.Show()

	return addForm
}

func (c *config) completeButton() func() {
	return func() {
		var TempData []Item

		for _, i := range c.TasksOnJSON {
			if c.TaskLabel.Text == i.Task {

				item := Item{
					Task:        i.Task,
					Description: i.Description,
					Done:        true,
					CreatedAt:   i.CreatedAt,
					CompletedAt: time.Now().Format("Mon 2 Jan 2006 15:04"),
				}

				TempData = append(TempData, item)
			} else {
				TempData = append(TempData, i)
			}
		}

		c.TasksOnJSON = TempData
		c.Store(TASKS_FILE)

		c.refreshPendings()
	}
}

func (c *config) deleteButton() func() {
	return func() {
		var TempData []Item

		for _, i := range c.TasksOnJSON {
			if c.TaskLabel.Text != i.Task {
				TempData = append(TempData, i)
			}
		}

		c.TasksOnJSON = TempData
		c.Store(TASKS_FILE)
		c.refreshPendings()
	}
}

func (c *config) refreshPendings() {
	c.Counter = c.CountPending()
	c.Pendings.Set(fmt.Sprintf("You have %d pending task(s)", c.Counter))
}
