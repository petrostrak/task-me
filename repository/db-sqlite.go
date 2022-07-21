package repository

import (
	"database/sql"
	"errors"
	"time"
)

type SQLiteRepository struct {
	Conn *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{Conn: db}
}

func (repo *SQLiteRepository) Migrate() error {
	query := `
	create table if not exists tasks(
		id integer primary key autoincrement,
		title text not null,
		description text,
		done numeric not null,
		created_at integer not null,
		completed_at integer not null);
	`

	_, err := repo.Conn.Exec(query)

	return err
}

func (repo *SQLiteRepository) InsertTask(t Task) (*Task, error) {
	stmt := `
	insert 
		into tasks (title, description, done, created_at, completed_at) 
		values (?, ?, ?, ?, ?)
	`

	result, err := repo.Conn.Exec(stmt, t.Title, t.Description, t.Done, t.CreatedAt.Unix(), t.CompletedAt.Unix())
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	t.ID = id

	return &t, nil
}

func (repo *SQLiteRepository) AllTasks() ([]Task, error) {
	query := `
	select
		id, title, description, done, created_at, completed_at
		from tasks order by created_at
	`

	rows, err := repo.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	var all []Task
	for rows.Next() {
		var t Task
		var unixTime int64
		err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&unixTime,
			&unixTime,
		)
		if err != nil {
			return nil, err
		}

		t.CreatedAt = time.Unix(unixTime, 0)
		t.CompletedAt = time.Unix(unixTime, 0)
		all = append(all, t)
	}

	return all, nil
}

func (repo *SQLiteRepository) GetTaskByID(id int) (*Task, error) {
	query := `
	select
		id, title, description, done, created_at, completed_at
		where id = ?
	`

	row := repo.Conn.QueryRow(query, id)

	var t Task
	var unixTime int64

	err := row.Scan(
		&t.ID,
		&t.Title,
		&t.Description,
		&unixTime,
		&unixTime,
	)
	if err != nil {
		return nil, err
	}

	t.CreatedAt = time.Unix(unixTime, 0)
	t.CompletedAt = time.Unix(unixTime, 0)

	return &t, nil
}

func (repo *SQLiteRepository) UpdateTask(id int64, updated Task) error {
	if id == 0 {
		return errors.New("update failed")
	}

	stmt := "update tasks set title = ?, description = ?, done = ?, completed_at = ? where id  = ?"

	result, err := repo.Conn.Exec(stmt, updated.Title, updated.Description, updated.Done, updated.CompletedAt.Unix(), id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("update failed")
	}

	return nil
}
