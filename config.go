package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	TaskList *widget.List
}

func (c *config) createMenuItems(win fyne.Window) {
	about := fyne.NewMenuItem("About", c.About(win))

	fileMenu := fyne.NewMenu("File", about)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)
}
