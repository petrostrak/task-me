package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const (
	TASKS_FILE = ".tasks.json"
)

func main() {
	a := app.NewWithID("app.petrostrak.taskMe.preferences")
	win := a.NewWindow("taskMe!")

	c := config{
		App:              a,
		Tasks:            make([]Item, 0),
		Counter:          0,
		Pendings:         binding.NewString(),
		TaskEntry:        widget.NewEntry(),
		DescriptionEntry: widget.NewEntry(),
		TaskLabels: TaskLabels{
			TaskLabel:        widget.NewLabel("Task"),
			DescriptionLabel: widget.NewLabel("Description"),
			CompletedLabel:   widget.NewLabel("Done?"),
			CreatedAtLabel:   widget.NewLabel("Created at"),
			CompletedAtLabel: widget.NewLabel("Completed at"),
		},
	}
	c.TaskLabels.TaskLabel.TextStyle = fyne.TextStyle{Bold: true}
	c.TaskEntry.SetPlaceHolder("Add a new task here")
	c.DescriptionEntry.SetPlaceHolder("Add description here")

	// open connection to DB
	db, err := c.connectSQL()
	if err != nil {
		log.Panic(err)
	}

	// create a DB repository
	c.setupDB(db)

	if err := c.Load(TASKS_FILE); err != nil {
		dialog.ShowError(err, win)
		os.Exit(1)
	}
	c.refreshPendings()

	// Define a welcome text centered
	text := c.WelcomeMessage()

	// Define the add button
	add, complete, delete, pending, _ := c.makeUI()

	table := c.tasks()

	// main menu
	c.createMenuItems(win)

	// Display content
	win.SetContent(container.NewHSplit(
		table,
		container.NewVBox(
			text, c.TaskLabel, c.DescriptionLabel,
			c.CompletedLabel, c.CreatedAtLabel, c.CompletedAtLabel,
			c.TaskEntry, c.DescriptionEntry, add, complete, delete,
			pending,
		),
	))

	win.Resize(fyne.NewSize(770, 410))
	win.CenterOnScreen()
	win.ShowAndRun()
}
