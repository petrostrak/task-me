package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	TASKS_FILE = ".tasks.json"
)

func main() {
	a := app.New()
	win := a.NewWindow("taskMe!")

	cfg := config{
		Tasks: make([]Item, 0),
	}

	if err := cfg.Load(TASKS_FILE); err != nil {
		os.Exit(1)
	}

	cfg.Pendings = cfg.CountPending()
	log.Println(cfg.Pendings)

	// main menu
	cfg.createMenuItems(win)

	// Define a welcome text centered
	text := cfg.WelcomeMessage()

	l_task := widget.NewLabel("Task")
	l_task.TextStyle = fyne.TextStyle{Bold: true}

	l_completed := widget.NewLabel("Done?")
	l_createdAt := widget.NewLabel("Created at")
	l_completedAt := widget.NewLabel("Completed at")

	e_task := widget.NewEntry()
	e_task.SetPlaceHolder("Add a new task here")

	pending := widget.NewLabel(fmt.Sprintf("You have %d pending task(s)", cfg.Pendings))
	pending.Alignment = fyne.TextAlignCenter

	// Define the add button
	addButton := widget.NewButton("Add a Task", func() {
		cfg.Add(e_task.Text)
		cfg.Store(TASKS_FILE)

		e_task.Text = ""
		e_task.Refresh()
		cfg.Pendings++
	})

	// Delete  button
	delete := widget.NewButton("Delete a Task", func() {
		var TempData []Item

		for _, i := range cfg.Tasks {
			if l_task.Text != i.Task {
				TempData = append(TempData, i)
			}
		}

		cfg.Tasks = TempData
		cfg.Store(TASKS_FILE)
	})

	// Render the list of tasks
	list := widget.NewList(
		func() int { return len(cfg.Tasks) },

		func() fyne.CanvasObject { return widget.NewLabel("") },

		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(cfg.Tasks[i].Task)
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		l_task.Text = cfg.Tasks[id].Task
		l_task.Refresh()
		if cfg.Tasks[id].Done {
			l_completed.Text = "Done!"
			l_completed.Refresh()
			l_completedAt.Text = cfg.Tasks[id].CompletedAt.Format("01 JAN 2006 15:04")
			l_completedAt.Refresh()
		} else {
			l_completed.Text = "Not done yet"
			l_completed.Refresh()
			l_completedAt.Text = "Pending..."
			l_completedAt.Refresh()
		}
		l_createdAt.Text = cfg.Tasks[id].CreatedAt.Format("01 JAN 2006 15:04")
		l_createdAt.Refresh()
	}

	// Complete  button
	complete := widget.NewButton("Complete a Task", func() {
		var TempData []Item

		for _, i := range cfg.Tasks {
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

		cfg.Tasks = TempData
		cfg.Store(TASKS_FILE)

		e_task.Text = ""
		e_task.Refresh()
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

	win.Resize(fyne.NewSize(600, 400))
	win.CenterOnScreen()
	win.ShowAndRun()
}
