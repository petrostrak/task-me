package main

import (
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/petrostrak/task-me/repository"
)

func (c *config) tasks() *fyne.Container {
	c.Tasks = c.getTaskSlice()
	c.TasksTable = c.getTasksTable()

	tasksContainer := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		container.NewAdaptiveGrid(1, c.TasksTable),
	)

	return tasksContainer
}

func (c *config) getTasksTable() *widget.Table {
	t := widget.NewTable(
		func() (int, int) {
			return len(c.Tasks), len(c.Tasks[0])
		},
		func() fyne.CanvasObject {
			return container.NewVBox(widget.NewLabel(""))
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == len(c.Tasks[0])-1 && i.Row != 0 {
				// last cell in row, put a button
				w := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
					dialog.ShowConfirm("Delete?", "", func(deleted bool) {
						if deleted {
							id, _ := strconv.Atoi(c.Tasks[i.Row][0].(string))
							err := c.DB.DeleteTask(int64(id))
							if err != nil {
								log.Println(err)
							}
						}

						// refresh the tasks table
						c.refreshTaskTable()
						// refresh pendings
						c.refreshPendings()
					}, c.MainWindow)
				})
				w.Importance = widget.HighImportance

				o.(*fyne.Container).Objects = []fyne.CanvasObject{w}
			} else if i.Col == len(c.Tasks[0])-2 && i.Row != 0 {
				w := widget.NewButtonWithIcon("Update", theme.ContentRedoIcon(), func() {
					id, _ := strconv.Atoi(c.Tasks[i.Row][0].(string))
					title := c.Tasks[i.Row][1].(string)
					desc := c.Tasks[i.Row][2].(string)
					done, _ := strconv.ParseBool(c.Tasks[i.Row][3].(string))
					created_at, _ := time.Parse("Mon 2 Jan 2006 15:04", c.Tasks[i.Row][4].(string))

					t := repository.Task{
						ID:          int64(id),
						Title:       title,
						Description: desc,
						Done:        done,
						CreatedAt:   created_at,
					}
					c.updateTaskDialog(t)
					// refresh the tasks table
					c.refreshTaskTable()
				})
				w.Importance = widget.MediumImportance

				o.(*fyne.Container).Objects = []fyne.CanvasObject{w}
			} else {
				// we are putting textual information
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(c.Tasks[i.Row][i.Col].(string)),
				}
			}
		},
	)

	colwidth := []float32{50, 170, 300, 105, 180, 110, 110}
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

	slice = append(slice, []any{"ID", "Title", "Description", "Done?", "Created at", "Update", "Delete"})

	for _, x := range tasks {
		var currentRow []any

		currentRow = append(currentRow, strconv.FormatInt(x.ID, 10))
		currentRow = append(currentRow, x.Title)
		currentRow = append(currentRow, x.Description)
		currentRow = append(currentRow, c.convertBool(x.Done))
		currentRow = append(currentRow, x.CreatedAt.Format("Mon 2 Jan 2006 15:04"))
		currentRow = append(currentRow, widget.NewButton("Update", func() {}))
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
	c.Tasks = c.getTaskSlice()
	c.TasksTable.Refresh()
}
