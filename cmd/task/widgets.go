package task

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (t *Tasks) AddButtonWidget(win fyne.Window) *widget.Button {
	return widget.NewButton("Add", func() {
		input := widget.NewEntry()
		input.SetPlaceHolder("Add a task")

		dialog.ShowCustom("What are you planning on doing?", "Close", container.NewVBox(
			input,
			widget.NewButton("Save", func() {
				t.Add(input.Text)
				fmt.Println(t)
			}),
		), win)
	})
}
