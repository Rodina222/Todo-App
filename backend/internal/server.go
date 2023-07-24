package server

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// ErrConnectToDB is returned when the app can't connect to the input database
var ErrConnectToDB = errors.New("failed to connect to the database")

// App is a struct that holds the database
type App struct {
	db *sql.DB
}

// Todo is a struct that holds one todo item
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// NewApp returns a new app that holds the database
func NewApp(db *sql.DB) *App {
	return &App{db}
}

// appRouter connects a pattern/path to a specific handler
func (app *App) AppRouter() error {

	router := gin.New()

	router.POST("/todos", app.CreateTodo)
	router.DELETE("/todos/:id", app.DeleteTodo)
	router.GET("/todos/:id", app.GetTodoByID)
	router.PUT("/todos/:id/completed", app.MarkAsCompleted)
	router.PATCH("/todos/:id", app.RenameTodo)

	err := http.ListenAndServe(":8080", router)

	return err
}

// ConnectToDB connects the app with the database
func ConnectToDB(db string) (*sql.DB, error) {

	database, err := sql.Open("sqlite3", db)

	if err != nil {
		return nil, fmt.Errorf("%v: %w", ErrConnectToDB, err)
	}

	fmt.Println(database.Begin())

	defer database.Close()

	return database, nil
}

// GetAllTodos returns all the todo list items
func (app *App) GetAllTodos(c *gin.Context) {

	var todos []Todo

	rows, err := app.db.Query("SELECT * FROM toDoList")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		todos = append(todos, todo)
	}

	c.JSON(http.StatusOK, todos)
}

// CreateTodo creates/adds a todo item to the database
func (app *App) CreateTodo(c *gin.Context) {

	var todo Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := app.db.Exec("INSERT INTO todos (title,completed) VALUES (?, ?)", todo.Title, todo.Completed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id, "title": todo.Title, "completed": todo.Completed})
}

// GetTodoByID returns the todo item by searching in the database by id
func (app *App) GetTodoByID(c *gin.Context) {

	var todo Todo
	id := c.Param("id")

	row := app.db.QueryRow("SELECT * FROM todos WHERE id = ?", id)

	err := row.Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DeleteTodo deletes a todo item given its id
func (app *App) DeleteTodo(c *gin.Context) {

	id := c.Param("id")

	_, err := app.db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Todo deleted"})
}

// MarkAsCompleted updates the state of todo item to complete using its id
func (app *App) MarkAsCompleted(c *gin.Context) {

	id := c.Param("id")

	_, err := app.db.Exec("UPDATE todos SET completed = true WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo marked as completed"})
}

// RenameTodo renames a todo item
func (app *App) RenameTodo(c *gin.Context) {

	id := c.Param("id")
	var todo Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := app.db.Exec("UPDATE todos SET title = ? WHERE id = ?", todo.Title, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Todo renamed"})
}
