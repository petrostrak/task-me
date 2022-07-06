package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
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

	// Define a welcome text centered
	text := canvas.NewText("Welcome to taskMe!", color.White)
	text.Alignment = fyne.TextAlignCenter

	// Initialize tasks
	tasks := task.Tasks{}

	// Define the add button
	addButton := tasks.AddButtonWidget(win, TODO_FILE)

	// Display a vertical box
	box := container.NewVBox(
		text,
		addButton,
	)

	// Display content
	win.SetContent(box)
	win.ShowAndRun()
}
