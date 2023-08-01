package internal

import (
	"database/sql"
	_ "embed"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var (
	// ErrConnectToDB is returned when the app can't connect to the input database
	ErrConnectToDB = errors.New("failed to connect to the database")

	// ErrCreateDBTable is returned when the app failed to create the database table
	ErrCreateDBTable = errors.New("failed to create the database table")
)

var (
	//go:embed sqlQueries/createTable.sql
	todosTable string

	//go:embed sqlQueries/createTodo.sql
	createTodo string

	//go:embed sqlQueries/deleteTodo.sql
	deleteTodo string

	//go:embed sqlQueries/getAllTodos.sql
	getAllTodos string

	//go:embed sqlQueries/getTodo.sql
	getTodo string

	//go:embed sqlQueries/updateTodo.sql
	updateTodo string
)

// DB is a struct that holds the database
type DB struct {
	db *sql.DB
}

// ConnectToDB connects the app to the database
func ConnectToDB(path string) (*sql.DB, error) {
	database, err := sql.Open("sqlite3", path)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrConnectToDB, err)
	}

	db := &DB{db: database}

	if err = db.CreateTable(); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCreateDBTable, err)
	}

	return db.db, nil
}

// CreateTable is responsible for creating the table of todos if the database is empty
func (db *DB) CreateTable() error {

	query, err := db.db.Prepare(todosTable)

	if err != nil {
		return err
	}

	query.Exec()

	return nil
}

func (db *DB) getAllTodos() (*sql.Rows, error) {

	rows, err := db.db.Query(getAllTodos)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (db *DB) getTodobyID(id string) *sql.Row {

	row := db.db.QueryRow(getTodo, id)

	return row

}

func (db *DB) createTodo(todo Todo) (sql.Result, error) {

	result, err := db.db.Exec(createTodo, todo.Title, todo.Completed)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db *DB) deleteTodo(id string) (sql.Result, error) {

	result, err := db.db.Exec(deleteTodo, id)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db *DB) updateTodo(todo Todo) (sql.Result, error) {

	result, err := db.db.Exec(updateTodo, todo.Title, todo.Completed, todo.ID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
