package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/petrostrak/todo-desktop-app-in-Go/cmd/task"
)

const (
	TODO_FILE = ".tasks.json"
)

func main() {
	taskMe := app.New()
	win := taskMe.NewWindow("taskMe!")

	// main menu
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Quit", func() { taskMe.Quit() }))

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowCustom("About", "Close", container.NewVBox(
				widget.NewLabel("Welcome to taskMe!, a simple todo Desktop app created in Go with Fyne."),
				widget.NewLabel("Version: v0.1"),
				widget.NewLabel("Author: Petros Trak"),
			), win)
		}))

	mainMenu := fyne.NewMainMenu(
		fileMenu,
		helpMenu,
	)
	win.SetMainMenu(mainMenu)
	win.Resize(fyne.NewSize(600, 400))

	// Initialize tasks and load tasks from file
	tasks := task.Tasks{}
	if err := tasks.Load(TODO_FILE); err != nil {
		os.Exit(1)
	}

	// Define a welcome text centered
	text := tasks.WelcomeMessage()

	// Render the list of tasks
	table := tasks.TableOfTasks(tasks)

	// Define the add button
	addButton := tasks.AddButtonWidget(win, TODO_FILE)

	// Display a vertical box
	box := container.New(
		&task.TableOfTasks{},
		text,
		table,
		addButton,
	)

	// Display content
	win.SetContent(box)
	win.ShowAndRun()
}
