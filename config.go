package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
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

func (c *config) makeUI() {
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			c.addTaskDialog()
		}),
	)

	pending := widget.NewLabelWithData(c.Pendings)
	pending.Alignment = fyne.TextAlignLeading

	table := c.tasks()

	// get app tabs
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Tasks", theme.InfoIcon(), table),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// add container to window
	finalContent := container.NewVBox(container.NewGridWithColumns(2, pending, toolbar), tabs)

	c.MainWindow.SetContent(finalContent)
}
