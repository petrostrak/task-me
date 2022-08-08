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

func (c *config) refreshPendings() {
	c.Counter = c.CountPending()
	c.Pendings.Set(fmt.Sprintf("You have %d pending task(s)", c.Counter))
}
