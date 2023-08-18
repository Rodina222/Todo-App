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

	if _, err = query.Exec(); err != nil {
		return err

	}

	if err = db.doesTableExists("todos"); err != nil {
		return err

	}

	return nil
}

func (db *DB) doesTableExists(tableName string) error {

	query := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s'", tableName)
	row := db.db.QueryRow(query)

	var name string
	err := row.Scan(&name)

	if err != nil {
		return err
	}

	return nil
}

// GetAllTodosDb returns all the todos in the database
func (db *DB) GetAllTodosDb() ([]Todo, error) {

	var todos []Todo

	rows, err := db.db.Query(getAllTodos)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

// GetTodobyIdDb returns the todo item by searching in the database by its id
func (db *DB) GetTodobyIdDb(id string) (Todo, error) {

	var todo Todo

	row := db.db.QueryRow(getTodo, id)

	if err := row.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
		return todo, err
	}

	return todo, nil

}

// CreateTodoDb inserts a todo in the database
func (db *DB) CreateTodoDb(todo Todo) (int, error) {

	result, err := db.db.Exec(createTodo, todo.Title, todo.Completed)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil

}

// DeleteTodoDb deletes a todo item given its id
func (db *DB) DeleteTodoDb(id string) (bool, error) {

	result, err := db.db.Exec(deleteTodo, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

// UpdateTodoDb updates the todo item using its id
func (db *DB) UpdateTodoDb(todo Todo) (bool, error) {
	result, err := db.db.Exec(updateTodo, todo.Title, todo.Completed, todo.ID)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
