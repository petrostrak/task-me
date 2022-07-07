package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (t *Tasks) FileMenu(taskMe fyne.App) *fyne.Menu {
	return fyne.NewMenu("File",
		fyne.NewMenuItem("Quit", func() { taskMe.Quit() }))
}

func (t *Tasks) HelpMenu(win fyne.Window) *fyne.Menu {
	return fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowCustom("About", "Close", container.NewVBox(
				widget.NewLabel("Welcome to taskMe!, a simple todo Desktop app written in Go with Fyne."),
				widget.NewLabel("Version: v2.0.0"),
				widget.NewLabel("Author: Petros Trak"),
			), win)
		}))
}

func (t *Tasks) WelcomeMessage() *canvas.Text {
	text := canvas.NewText("Welcome to taskMe!", color.White)
	text.Alignment = fyne.TextAlignCenter
	text.Resize(fyne.NewSize(600, 50))

	return text
}

func (t *Tasks) PendingTasks() *canvas.Text {
	text := canvas.NewText(fmt.Sprintf("You have %d pending task(s)", t.CountPending()), color.White)
	text.Alignment = fyne.TextAlignCenter
	text.Resize(fyne.NewSize(600, 50))

	return text
}
