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
				c.refreshPendings()
			}
		}, c.MainWindow)

	// size and show the dialog
	addForm.Resize(fyne.Size{Width: 400})
	addForm.Show()

	return addForm
}

func (c *config) updateTaskDialog(id int) dialog.Dialog {
	var done string
	updateLableEntry := widget.NewEntry()
	updateDescriptionEntry := widget.NewEntry()
	updateDoneEntry := widget.NewSelect([]string{"Done", "Not yet done"}, func(s string) {
		done = s
	})
	updateDoneEntry.Selected = done

	c.AddTasksLableEntry = updateLableEntry
	c.AddTasksDescriptionEntry = updateDescriptionEntry
	c.UpdateDoneEntry = updateDoneEntry

	updateForm := dialog.NewForm(
		"Update Task",
		"Update",
		"Cancel",
		[]*widget.FormItem{
			{Text: "Title of task", Widget: updateLableEntry},
			{Text: "Description", Widget: updateDescriptionEntry},
			{Text: "Done?", Widget: updateDoneEntry},
		},
		func(valid bool) {
			if valid {
				t, _ := c.DB.GetTaskByID(id)
				title := updateLableEntry.Text
				description := updateDescriptionEntry.Text
				var isDone bool
				if done == "Done" {
					isDone = true
				} else {
					isDone = false
				}

				t.Title = title
				t.Description = description
				t.Done = isDone
				t.CompletedAt = time.Now()

				err := c.DB.UpdateTask(int64(id), *t)
				if err != nil {
					log.Println(err)
				}

				c.refreshTaskTable()
				c.refreshPendings()
			}
		}, c.MainWindow)

	// size and show the dialog
	updateForm.Resize(fyne.Size{Width: 400})
	updateForm.Show()

	return updateForm
}

func (c *config) refreshPendings() {
	c.Counter = c.CountPending()
	c.Pendings.Set(fmt.Sprintf("You have %d pending task(s)", c.Counter))
}
