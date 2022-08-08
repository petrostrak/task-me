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

	c := config{
		App:         a,
		TasksOnJSON: make([]Item, 0),
		Counter:     0,
		Pendings:    binding.NewString(),
		TaskLabels: TaskLabels{
			TaskLabel:        widget.NewLabel("Task"),
			DescriptionLabel: widget.NewLabel("Description"),
			CompletedLabel:   widget.NewLabel("Done?"),
			CreatedAtLabel:   widget.NewLabel("Created at"),
			CompletedAtLabel: widget.NewLabel("Completed at"),
		},
		MainWindow: a.NewWindow("taskMe!"),
	}
	c.TaskLabels.TaskLabel.TextStyle = fyne.TextStyle{Bold: true}

	// open connection to DB
	db, err := c.connectSQL()
	if err != nil {
		log.Panic(err)
	}

	// create a DB repository
	c.setupDB(db)

	if err := c.Load(TASKS_FILE); err != nil {
		dialog.ShowError(err, c.MainWindow)
		os.Exit(1)
	}
	c.refreshPendings()

	// Define a welcome text centered
	text := c.WelcomeMessage()

	// Define the add button
	add, complete, pending, _ := c.makeUI()

	table := c.tasks()

	// main menu
	c.createMenuItems(c.MainWindow)

	// Display content
	c.MainWindow.SetContent(container.NewGridWithColumns(2,
		table,
		container.NewVBox(
			text, c.TaskLabel, c.DescriptionLabel,
			c.CompletedLabel, c.CreatedAtLabel, c.CompletedAtLabel,
			add, complete,
			pending,
		),
	))

	c.MainWindow.Resize(fyne.NewSize(770, 410))
	c.MainWindow.CenterOnScreen()
	c.MainWindow.ShowAndRun()
}
