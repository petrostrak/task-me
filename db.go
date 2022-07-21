package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/glebarez/go-sqlite"
	"github.com/petrostrak/task-me/repository"
)

func (c *config) connectSQL() (*sql.DB, error) {
	path := ""

	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	} else {
		path = c.App.Storage().RootURI().Path() + "/sql.db"
		log.Println("db in:", path)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (c *config) setupDB(db *sql.DB) {
	c.DB = repository.NewSQLiteRepository(db)

	err := c.DB.Migrate()
	if err != nil {
		log.Panicln(err)
	}
}
