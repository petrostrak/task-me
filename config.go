package main

import (
	"fyne.io/fyne/v2"
)

type config struct {
	Tasks    []Item
	Pendings int
}

func (c *config) createMenuItems(win fyne.Window) {
	about := fyne.NewMenuItem("About", c.About(win))

	fileMenu := fyne.NewMenu("File", about)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)
}
