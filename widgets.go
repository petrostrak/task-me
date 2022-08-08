package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func (c *config) PendingTasks() *canvas.Text {
	text := canvas.NewText(fmt.Sprintf("You have %d pending task(s)", c.CountPending()), color.White)
	text.Alignment = fyne.TextAlignCenter
	text.Resize(fyne.NewSize(600, 50))

	return text
}
