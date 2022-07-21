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

func TestSQLiteRepository_GetTaskByID(t *testing.T) {
	tsk, err := testRepo.GetTaskByID(1)
	if err != nil {
		t.Error("get by id failed:", err)
	}

	if tsk.Description != "description" {
		t.Errorf("wrong task description; expected 'description' but got %s", tsk.Description)
	}

	_, err = testRepo.GetTaskByID(2)
	if err == nil {
		t.Error("get values from non-existent id")
	}
}
