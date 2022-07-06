package task

import (
	"strconv"

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
				t.Store(filename)
			}),
		), win)
	})
}

func (t *Tasks) TableOfTasks(tasks Tasks) *widget.Table {
	var data = [][]string{}
	header := []string{"#", "Task", "Done?", "Created at", "Completed at"}
	data = append(data, header)

	for idx, x := range *t {
		idx++
		done := "no"
		completed := "Not yet completed"
		if x.Done {
			done = "yes"
			completed = x.CompletedAt.Format("01 JAN 2006 15:04")
		}

		var item []string
		item = append(item, strconv.Itoa(idx))
		item = append(item, x.Task)
		item = append(item, done)
		item = append(item, x.CreatedAt.Format("01 JAN 2006 15:04"))
		item = append(item, completed)

		data = append(data, item)
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("tasks")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		},
	)

	table.SetColumnWidth(0, 30)
	table.SetColumnWidth(1, 190)
	table.SetColumnWidth(2, 60)
	table.SetColumnWidth(3, 150)
	table.SetColumnWidth(4, 150)

	return table
}
