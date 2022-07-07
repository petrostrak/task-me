package main

import (
	"fmt"
	"image/color"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	TASKS_FILE = ".tasks.json"
)

func main() {
	taskMe := app.New()
	win := taskMe.NewWindow("taskMe!")

	// Initialize tasks and load tasks from file
	tasks := Tasks{}
	if err := tasks.Load(TASKS_FILE); err != nil {
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

	l_task := widget.NewLabel("Task")
	l_task.TextStyle = fyne.TextStyle{Bold: true}

	l_completed := widget.NewLabel("Done?")
	l_createdAt := widget.NewLabel("Created at")
	l_completedAt := widget.NewLabel("Completed at")

	e_task := widget.NewEntry()
	e_task.SetPlaceHolder("Add a new task here")

	pending := canvas.NewText(fmt.Sprintf("You have %d pending task(s)", tasks.CountPending()), color.White)
	pending.Alignment = fyne.TextAlignCenter

	// Define the add button
	addButton := widget.NewButton("Add a Task", func() {
		tasks.Add(e_task.Text)
		tasks.Store(TASKS_FILE)

		e_task.Text = ""
		e_task.Refresh()

		pending.Refresh()
	})

	// Delete  button
	delete := widget.NewButton("Delete a Task", func() {
		var TempData []Item

		for _, i := range tasks {
			if l_task.Text != i.Task {
				TempData = append(TempData, i)
			}
		}

		tasks = TempData
		tasks.Store(TASKS_FILE)

		pending.Refresh()
	})

	// Render the list of tasks
	list := widget.NewList(
		func() int { return len(tasks) },

		func() fyne.CanvasObject { return widget.NewLabel("") },

		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(tasks[i].Task)
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		l_task.Text = tasks[id].Task
		l_task.Refresh()
		if tasks[id].Done {
			l_completed.Text = "Done!"
			l_completed.Refresh()
			l_completedAt.Text = tasks[id].CompletedAt.Format("01 JAN 2006 15:04")
			l_completedAt.Refresh()
		} else {
			l_completed.Text = "Not done yet"
			l_completed.Refresh()
			l_completedAt.Text = "Pending..."
			l_completedAt.Refresh()
		}
		l_createdAt.Text = tasks[id].CreatedAt.Format("01 JAN 2006 15:04")
		l_createdAt.Refresh()
	}

	// Complete  button
	complete := widget.NewButton("Complete a Task", func() {
		var TempData []Item

		for _, i := range tasks {
			if l_task.Text == i.Task {

				item := Item{
					Task:        i.Task,
					Done:        true,
					CreatedAt:   i.CreatedAt,
					CompletedAt: time.Now(),
				}

				TempData = append(TempData, item)
			} else {
				TempData = append(TempData, i)
			}
		}

		tasks = TempData
		tasks.Store(TASKS_FILE)

		e_task.Text = ""
		e_task.Refresh()

		list.Refresh()
		pending.Refresh()
	})

	// Display content
	win.SetContent(container.NewHSplit(
		list,
		container.NewVBox(
			text, l_task, l_completed, l_createdAt, l_completedAt,
			e_task, addButton, complete, delete,
			pending,
		),
	))
	win.ShowAndRun()
}
