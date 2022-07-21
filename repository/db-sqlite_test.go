package repository

import (
	"testing"
	"time"
)

func TestSQLiteRepository_Migrate(t *testing.T) {
	err := testRepo.Migrate()

	if err != nil {
		t.Error("migrate failed:", err)
	}
}

func TestSQLiteRepository_InsertTask(t *testing.T) {
	tsk := Task{
		Title:       "title",
		Description: "description",
		Done:        true,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}

	result, err := testRepo.InsertTask(tsk)
	if err != nil {
		t.Error("insert failed:", err)
	}

	if result.ID <= 0 {
		t.Error("invalid id sent back:", result.ID)
	}
}
