package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/petrostrak/todo-desktop-app-in-Go/cmd/task"
)

const (
	TODO_FILE = ".tasks.json"
)

func main() {
	taskMe := app.New()
	win := taskMe.NewWindow("taskMe!")

	// Initialize tasks and load tasks from file
	tasks := task.Tasks{}
	if err := tasks.Load(TODO_FILE); err != nil {
		os.Exit(1)
	}

	// main menu
	fileMenu := tasks.FileMenu(taskMe)

	helpMenu := tasks.HelpMenu(win)

	mainMenu := fyne.NewMainMenu(
		fileMenu,
		helpMenu,
	)
	win.SetMainMenu(mainMenu)
	win.Resize(fyne.NewSize(600, 400))

	// Define a welcome text centered
	text := tasks.WelcomeMessage()

	// Render the list of tasks
	table := tasks.TableOfTasks(tasks)

	// Define the add button
	addButton := tasks.AddButtonWidget(win, TODO_FILE, table)

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
