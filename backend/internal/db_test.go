package internal

import (
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectToDB(t *testing.T) {

	tempDir := t.TempDir()

	dbPath := filepath.Join(tempDir, "database.db")

	_, err := ConnectToDB(dbPath)
	assert.NoError(t, err)

}

func TestCreateTable(t *testing.T) {

	tempDir := t.TempDir()

	dbPath := filepath.Join(tempDir, "database.db")

	database, err := ConnectToDB(dbPath)
	assert.NoError(t, err)

	db := &DB{db: database}

	err = db.CreateTable()
	assert.NoError(t, err)

}

func TestCreateTodoDb(t *testing.T) {

	tempDir := t.TempDir()

	dbPath := filepath.Join(tempDir, "database.db")

	database, err := ConnectToDB(dbPath)
	assert.NoError(t, err)

	db := &DB{db: database}

	err = db.CreateTable()
	assert.NoError(t, err)

	newTodo := Todo{
		ID:        1,
		Title:     "Todo",
		Completed: false,
	}

	_, err = db.CreateTodoDb(newTodo)
	assert.NoError(t, err)

}

func TestDeleteTodoDb(t *testing.T) {

	tempDir := t.TempDir()

	dbPath := filepath.Join(tempDir, "database.db")

	database, err := ConnectToDB(dbPath)
	assert.NoError(t, err)

	db := &DB{db: database}

	err = db.CreateTable()
	assert.NoError(t, err)

	newTodo := Todo{
		ID:        1,
		Title:     "Todo",
		Completed: false,
	}

	id, err := db.CreateTodoDb(newTodo)
	assert.NoError(t, err)

	deleted, err := db.DeleteTodoDb(strconv.Itoa(id))
	assert.NoError(t, err)
	assert.True(t, deleted)

}

func TestUpdateTodoDb(t *testing.T) {

	tempDir := t.TempDir()

	dbPath := filepath.Join(tempDir, "database.db")

	database, err := ConnectToDB(dbPath)
	assert.NoError(t, err)

	db := &DB{db: database}

	err = db.CreateTable()
	assert.NoError(t, err)

	newTodo := Todo{
		ID:        1,
		Title:     "Todo",
		Completed: false,
	}

	_, err = db.CreateTodoDb(newTodo)
	assert.NoError(t, err)

	updatedTodo := Todo{
		ID:        1,
		Title:     "Todo",
		Completed: true,
	}

	updated, err := db.UpdateTodoDb(updatedTodo)
	assert.NoError(t, err)
	assert.True(t, updated)

}

func TestGetAllTodosDb(t *testing.T) {

	tempDir := t.TempDir()

	dbPath := filepath.Join(tempDir, "database.db")

	database, err := ConnectToDB(dbPath)
	assert.NoError(t, err)

	db := &DB{db: database}

	err = db.CreateTable()
	assert.NoError(t, err)

	newTodo := Todo{
		ID:        1,
		Title:     "Todo1",
		Completed: false,
	}

	_, err = db.CreateTodoDb(newTodo)
	assert.NoError(t, err)

	newTodo = Todo{
		ID:        2,
		Title:     "Todo2",
		Completed: false,
	}

	_, err = db.CreateTodoDb(newTodo)
	assert.NoError(t, err)

	_, err = db.GetAllTodosDb()
	assert.NoError(t, err)

}

func TestGetTodoByIdDb(t *testing.T) {

	tempDir := t.TempDir()

	dbPath := filepath.Join(tempDir, "database.db")

	database, err := ConnectToDB(dbPath)
	assert.NoError(t, err)

	db := &DB{db: database}

	err = db.CreateTable()
	assert.NoError(t, err)

	newTodo := Todo{
		ID:        1,
		Title:     "Todo",
		Completed: false,
	}

	id, err := db.CreateTodoDb(newTodo)
	assert.NoError(t, err)

	_, err = db.GetTodobyIdDb(strconv.Itoa(id))
	assert.NoError(t, err)

}
