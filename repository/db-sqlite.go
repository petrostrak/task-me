package repository

import "database/sql"

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
