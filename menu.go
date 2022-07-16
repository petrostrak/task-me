package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (c *config) About(win fyne.Window) func() {
	return func() {
		dialog.ShowCustom("About", "Close", container.NewVBox(
			widget.NewLabel("Welcome to taskMe!, a simple todo Desktop app written in Go with Fyne."),
			widget.NewLabel("Version: v2.0.1"),
			widget.NewLabel("Author: Petros Trak"),
		), win)
	}
}
