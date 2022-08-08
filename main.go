package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

const (
	TASKS_FILE = ".tasks.json"
)

func main() {
	a := app.NewWithID("app.petrostrak.taskMe.preferences")

	c := config{
		App:        a,
		Counter:    0,
		Pendings:   binding.NewString(),
		MainWindow: a.NewWindow("taskMe!"),
	}

	// open connection to DB
	db, err := c.connectSQL()
	if err != nil {
		log.Panic(err)
	}

	// create a DB repository
	c.setupDB(db)

	c.refreshPendings()

	// Define a welcome text centered
	// text := c.WelcomeMessage()

	// Define the add button
	add, pending := c.makeUI()

	table := c.tasks()

	// main menu
	c.createMenuItems(c.MainWindow)

	// Display content
	c.MainWindow.SetContent(container.NewGridWithRows(2,
		table,
		container.NewVBox(add, pending),
	))

	c.MainWindow.Resize(fyne.NewSize(930, 410))
	c.MainWindow.CenterOnScreen()
	c.MainWindow.ShowAndRun()
}
