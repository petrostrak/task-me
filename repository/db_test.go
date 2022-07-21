package repository

import "time"

type TestRepository struct{}

func NewTestRepository() *TestRepository {
	return &TestRepository{}
}

func (repo *TestRepository) Migrate() error {
	return nil
}

func (repo *TestRepository) InsertHolding(t Task) (*Task, error) {
	return &t, nil
}

func (repo *TestRepository) AllTask() ([]Task, error) {
	var all []Task

	return all, nil
}

func (repo *TestRepository) GetHoldingByID(id int) (*Task, error) {
	t := Task{
		Title:       "title",
		Description: "description",
		Done:        true,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}

	return &t, nil
}

func (repo *TestRepository) UpdateHolding(id int64, updated Task) error {
	return nil
}

func (repo *TestRepository) DeleteHolding(id int64) error {
	return nil
}
