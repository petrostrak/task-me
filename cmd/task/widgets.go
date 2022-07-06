package task

import (
	"errors"
	"image/color"
	"strconv"

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
				widget.NewLabel("Version: v1.0.0"),
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

func (t *Tasks) AddButtonWidget(win fyne.Window, filename string) *widget.Button {
	button := widget.NewButton("Add a Task", func() {
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

	button.Resize(fyne.NewSize(592, 50))

	return button
}

func (t *Tasks) CompleteTask(win fyne.Window, filename string) *widget.Button {
	button := widget.NewButton("Complete a Task", func() {
		input := widget.NewEntry()
		input.SetPlaceHolder("# of task to mark as complete")

		dialog.ShowCustom("Choose a task to mark as complete!", "Close", container.NewVBox(
			input,
			widget.NewButton("Complete", func() {
				i, err := strconv.Atoi(input.Text)
				if err != nil {
					input.SetValidationError(errors.New("please give a number of task to mark as complete"))
				}

				t.Complete(i)
				t.Store(filename)
			}),
		), win)
	})

	button.Resize(fyne.NewSize(592, 50))

	return button
}

func (t *Tasks) DeleteTask(win fyne.Window, filename string) *widget.Button {
	button := widget.NewButton("Delete a Task", func() {
		input := widget.NewEntry()
		input.SetPlaceHolder("# of task to delete")

		dialog.ShowCustom("Choose a task to delete!", "Close", container.NewVBox(
			input,
			widget.NewButton("Delete", func() {
				i, err := strconv.Atoi(input.Text)
				if err != nil {
					input.SetValidationError(errors.New("please give a number of task to delete"))
				}

				t.Delete(i)
				t.Store(filename)
			}),
		), win)
	})

	button.Resize(fyne.NewSize(592, 50))

	return button
}

type TableOfTasks struct{}

func (t *TableOfTasks) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.Size()

		w += 0
		h += childSize.Height
	}
	return fyne.NewSize(w, h)
}

func (t *TableOfTasks) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, 0)

	for _, o := range objects {
		size := o.Size()
		o.Resize(size)
		o.Move(pos)

		pos = pos.Add(fyne.NewPos(0, size.Height))
	}
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

	table.Resize(fyne.NewSize(600, 400))

	return table
}
