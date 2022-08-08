package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/petrostrak/task-me/repository"
)

type config struct {
	App                      fyne.App
	Tasks                    [][]any
	TasksTable               *widget.Table
	MainWindow               fyne.Window
	Counter                  int
	Pendings                 binding.String
	DB                       repository.Repository
	AddTasksLableEntry       *widget.Entry
	AddTasksDescriptionEntry *widget.Entry
	UpdateDoneEntry          *widget.Select
}

func (c *config) createMenuItems(win fyne.Window) {
	about := fyne.NewMenuItem("About", c.About(win))

	fileMenu := fyne.NewMenu("File", about)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)
}

func (c *config) makeUI() (add *widget.Button, pending *widget.Label, table *fyne.Container) {
	add = widget.NewButton("Add a Task", func() {
		c.addTaskDialog()
	})

	// update = widget.NewButton("Update a Task", func() {
	// 	c.updateTaskDialog()
	// })

	pending = widget.NewLabelWithData(c.Pendings)
	pending.Alignment = fyne.TextAlignCenter

	table = c.tasks()

	return
}
