package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

const (
	TASKS_FILE = ".tasks.json"
)

func main() {
	a := app.New()
	win := a.NewWindow("taskMe!")

	cfg := config{
		Tasks:     make([]Item, 0),
		TaskEntry: widget.NewEntry(),
		TaskLabels: TaskLabels{
			TaskLabel:        widget.NewLabel("Task"),
			CompletedLabel:   widget.NewLabel("Done?"),
			CreatedAtLabel:   widget.NewLabel("Created at"),
			CompletedAtLabel: widget.NewLabel("Completed at"),
		},
	}
	cfg.TaskLabels.TaskLabel.TextStyle = fyne.TextStyle{Bold: true}
	cfg.TaskEntry.SetPlaceHolder("Add a new task here")

	if err := cfg.Load(TASKS_FILE); err != nil {
		os.Exit(1)
	}

	cfg.Pendings = cfg.CountPending()

	// Define a welcome text centered
	text := cfg.WelcomeMessage()

	data := binding.NewString()
	s := fmt.Sprintf("You have %d pending task(s)", cfg.Pendings)
	data.Set(s)

	pending := widget.NewLabelWithData(data)
	pending.Alignment = fyne.TextAlignCenter

	// Define the add button
	add, complete, delete, list := cfg.makeUI()

	// main menu
	cfg.createMenuItems(win)

	// Display content
	win.SetContent(container.NewHSplit(
		list,
		container.NewVBox(
			text, cfg.TaskLabel, cfg.CompletedLabel, cfg.CreatedAtLabel, cfg.CompletedAtLabel,
			cfg.TaskEntry, add, complete, delete,
			pending,
		),
	))

	win.Resize(fyne.NewSize(600, 400))
	win.CenterOnScreen()
	win.ShowAndRun()
}
