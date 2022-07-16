package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const (
	TASKS_FILE = ".tasks.json"
)

func main() {
	a := app.New()
	win := a.NewWindow("taskMe!")

	c := config{
		Tasks:     make([]Item, 0),
		Pendings:  0,
		TaskEntry: widget.NewEntry(),
		TaskLabels: TaskLabels{
			TaskLabel:        widget.NewLabel("Task"),
			CompletedLabel:   widget.NewLabel("Done?"),
			CreatedAtLabel:   widget.NewLabel("Created at"),
			CompletedAtLabel: widget.NewLabel("Completed at"),
		},
	}
	c.TaskLabels.TaskLabel.TextStyle = fyne.TextStyle{Bold: true}
	c.TaskEntry.SetPlaceHolder("Add a new task here")

	if err := c.Load(TASKS_FILE); err != nil {
		dialog.ShowError(err, win)
		os.Exit(1)
	}
	c.Pendings = c.CountPending()

	// Define a welcome text centered
	text := c.WelcomeMessage()

	// Define the add button
	add, complete, delete, pending, list := c.makeUI()

	// main menu
	c.createMenuItems(win)

	// Display content
	win.SetContent(container.NewHSplit(
		list,
		container.NewVBox(
			text, c.TaskLabel, c.CompletedLabel, c.CreatedAtLabel, c.CompletedAtLabel,
			c.TaskEntry, add, complete, delete,
			pending,
		),
	))

	win.Resize(fyne.NewSize(600, 400))
	win.CenterOnScreen()
	win.ShowAndRun()
}
