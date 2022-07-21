package repository

import "time"

type Repository interface {
	Migrate() error
	InsertTask(Task) (*Task, error)
	AllTasks() ([]Task, error)
	GetTaskByID(int) (*Task, error)
	UpdateTask(int64, Task) error
	DeleteTask(int64) error
}

type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}
