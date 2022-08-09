package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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

	// main menu
	c.createMenuItems(c.MainWindow)

	c.MainWindow.Resize(fyne.NewSize(1040, 410))
	c.MainWindow.CenterOnScreen()
	c.MainWindow.SetFixedSize(true)
	c.MainWindow.SetMaster()

	c.makeUI()
	c.MainWindow.ShowAndRun()
}
