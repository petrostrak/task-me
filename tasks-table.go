package main

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/petrostrak/task-me/repository"
)

func (c *config) getTasksTable() *widget.Table {
	data := c.getTaskSlice()
	c.TaskTable = data

	t := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return container.NewVBox(widget.NewLabel(""))
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == len(data[0])-1 && i.Row != 0 {
				// last cell in row, put a button
				w := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
					dialog.ShowConfirm("Delete?", "Cancel", func(deleted bool) {
						id, _ := strconv.Atoi(data[i.Row][0].(string))
						err := c.DB.DeleteTask(int64(id))
						if err != nil {
							log.Println(err)
						}

						// refresh the tasks table
						c.refreshTaskTable()
					}, c.MainWindow)
				})
				w.Importance = widget.HighImportance

				o.(*fyne.Container).Objects = []fyne.CanvasObject{w}
			} else {
				// we are putting textual information
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(data[i.Row][i.Col].(string)),
				}
			}
		},
	)

	colwidth := []float32{50, 600, 110}
	for i := 0; i < len(colwidth); i++ {
		t.SetColumnWidth(i, colwidth[i])
	}

	return t
}

func (c *config) getTaskSlice() [][]any {
	var slice [][]any

	tasks, err := c.currentTasks()
	if err != nil {
		log.Println(err)
	}

	slice = append(slice, []any{"ID", "Title", "Complete", "Delete"})

	for _, x := range tasks {
		var currentRow []any

		currentRow = append(currentRow, strconv.FormatInt(x.ID, 10))
		currentRow = append(currentRow, x.Title)
		currentRow = append(currentRow, widget.NewButton("Delete", func() {}))

		slice = append(slice, currentRow)
	}

	return slice
}

func (c *config) currentTasks() ([]repository.Task, error) {
	tasks, err := c.DB.AllTasks()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (c *config) refreshTaskTable() {
	c.TaskTable = c.getTaskSlice()
	c.Table.Refresh()
}
