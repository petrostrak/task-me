package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/petrostrak/task-me/repository"
)

type config struct {
	App              fyne.App
	TasksOnJSON      []Item
	Tasks            [][]any
	TasksTable       *widget.Table
	MainWindow       fyne.Window
	Counter          int
	Pendings         binding.String
	TaskEntry        *widget.Entry
	DescriptionEntry *widget.Entry
	DB               repository.Repository
	TaskLabels
}

type TaskLabels struct {
	TaskLabel        *widget.Label
	DescriptionLabel *widget.Label
	CompletedLabel   *widget.Label
	CreatedAtLabel   *widget.Label
	CompletedAtLabel *widget.Label
}

func (c *config) createMenuItems(win fyne.Window) {
	about := fyne.NewMenuItem("About", c.About(win))

	fileMenu := fyne.NewMenu("File", about)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)
}

func (c *config) makeUI() (add, complete, delete *widget.Button, pending *widget.Label, list *widget.List) {
	add = widget.NewButton("Add a Task", c.addButton())

	complete = widget.NewButton("Complete a Task", c.completeButton())

	delete = widget.NewButton("Delete a Task", c.deleteButton())

	pending = widget.NewLabelWithData(c.Pendings)
	pending.Alignment = fyne.TextAlignCenter

	list = widget.NewList(
		func() int { return len(c.TasksOnJSON) },

		func() fyne.CanvasObject { return widget.NewLabel("") },

		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(c.TasksOnJSON[i].Task)
		},
	)

	list.OnSelected = c.onSelect

	return
}

func (c *config) onSelect(id widget.ListItemID) {
	c.TaskLabel.Text = c.TasksOnJSON[id].Task
	c.TaskLabel.Refresh()
	if c.TasksOnJSON[id].Done {
		c.CompletedLabel.Text = "Done!"
		c.CompletedLabel.Refresh()
		c.CompletedAtLabel.Text = c.TasksOnJSON[id].CompletedAt
		c.CompletedAtLabel.Refresh()
	} else {
		c.CompletedLabel.Text = "Not done yet"
		c.CompletedLabel.Refresh()
		c.CompletedAtLabel.Text = "Pending..."
		c.CompletedAtLabel.Refresh()
	}
	c.CreatedAtLabel.Text = c.TasksOnJSON[id].CreatedAt
	c.CreatedAtLabel.Refresh()
	c.DescriptionLabel.Text = c.TasksOnJSON[id].Description
	c.DescriptionLabel.Refresh()
}
