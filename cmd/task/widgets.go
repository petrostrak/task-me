package task

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (t *Tasks) AddButtonWidget(win fyne.Window, filename string) *widget.Button {
	return widget.NewButton("Add", func() {
		input := widget.NewEntry()
		input.SetPlaceHolder("Add a task")

		dialog.ShowCustom("What are you planning on doing?", "Close", container.NewVBox(
			input,
			widget.NewButton("Save", func() {
				t.Add(input.Text)
				fmt.Println(t)
				t.Store(filename)
			}),
		), win)
	})
}

func (t *Tasks) ListOfTasks(tasks Tasks) *widget.List {
	return widget.NewList(
		func() int {
			return len(tasks)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("tasks")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(tasks[i].Task)
		},
	)
}
