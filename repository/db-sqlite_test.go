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

func TestSQLiteRepository_AllTasks(t *testing.T) {
	tsks, err := testRepo.AllTasks()
	if err != nil {
		t.Error("get all tasks failed:", err)
	}

	if len(tsks) != 1 {
		t.Errorf("wrong number of rows returned; expected 1, but got %d", len(tsks))
	}
}

func TestSQLiteRepository_UpdateTask(t *testing.T) {
	tsk, err := testRepo.GetTaskByID(1)
	if err != nil {
		t.Error(err)
	}

	tsk.Description = "need more testing"
	tsk.Done = false

	err = testRepo.UpdateTask(1, *tsk)
	if err != nil {
		t.Error("update failed:", err)
	}
}

func TestSQLiteRepository_DeleteTask(t *testing.T) {
	err := testRepo.DeleteTask(1)
	if err != nil {
		t.Error("failed to delete task:", err)
	}

	err = testRepo.DeleteTask(2)
	if err == nil {
		t.Error("no error when trying to delete non-existent record!")
	}
}
