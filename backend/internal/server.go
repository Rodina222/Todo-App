package internal

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	_ "github.com/codescalersinternships/ToDoApp-Rodina/backend/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swagFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// App is a struct that holds a struct of DB
type App struct {
	db DB
}

// Todo is a struct that holds one todo item
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// ErrorResponse is a struct that holds an error message
type ErrorResponse struct {
	Error string `json:"error"`
}

// NewApp returns a new app that holds the database
func NewApp(db *sql.DB) *App {

	appDB := DB{db: db}
	return &App{db: appDB}

}

// Run calls the internal appRouter method to create the app router and start the server
func (app *App) Run(ginMode string) error {

	if err := app.appRouter(ginMode); err != nil {
		return err
	}

	return nil
}

func (app *App) appRouter(ginMode string) error {

	gin.SetMode(ginMode)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.GET("/docs/*any", ginSwagger.WrapHandler(swagFiles.Handler))
	router.POST("/todos", app.CreateTodo)
	router.DELETE("/todos/:id", app.DeleteTodo)
	router.GET("/todos/:id", app.GetTodoByID)
	router.GET("/todos", app.GetAllTodos)
	router.PUT("/todos/:id", app.UpdateTodo)

	err := router.Run(":8096")

	return err
}

// GetAllTodos returns all the todos in the database
// @Summary Gets all todos
// @Description Gets all todos in the database
// @Tags todos
// @Produce json
// @Success 200 {array} Todo
// @Failure 500 {object} ErrorResponse "failed to get all todos"
// @Router /todos [get]
func (app *App) GetAllTodos(c *gin.Context) {

	error := ErrorResponse{"failed to get all todos"}

	todos, err := app.db.GetAllTodosDb()

	if err != nil {
		log.Println("error:", err)
		c.JSON(http.StatusInternalServerError, error)
		return
	}

	c.JSON(http.StatusOK, todos)
}

// CreateTodo inserts a todo in the database
// @Summary Adds a todo
// @Description Adds a new todo to the database
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body Todo true "New todo object"
// @Success 200 {object} Todo
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "failed to create todo"
// @Router /todos [post]
func (app *App) CreateTodo(c *gin.Context) {

	var todo Todo
	error := ErrorResponse{"failed to create todo"}

	if err := c.ShouldBindJSON(&todo); err != nil {
		log.Println("error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create todo"})
		return
	}

	id, err := app.db.CreateTodoDb(todo)

	if err != nil {
		log.Println("error:", err)
		c.JSON(http.StatusInternalServerError, error)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id, "title": todo.Title, "completed": todo.Completed})
}

// GetTodoByID returns the todo item by searching in the database by id
// @Summary Get a todo
// @Description Get a specific todo by id
// @Tags todos
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} Todo
// @Failure 404 {object} ErrorResponse "Todo not found"
// @Failure 500 {object} ErrorResponse "failed to get todo"
// @Router /todos/{id} [get]
func (app *App) GetTodoByID(c *gin.Context) {

	id := c.Param("id")

	error := ErrorResponse{"failed to get todo"}

	todo, err := app.db.GetTodobyIdDb(id)

	if err == sql.ErrNoRows {
		log.Println("error:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return

	}

	if err != nil {
		log.Println("error:", err)
		c.JSON(http.StatusInternalServerError, error)
		return

	}

	c.JSON(http.StatusOK, todo)
}

// DeleteTodo deletes a todo item given its id
// @Summary Deletes a todo
// @Description Deletes a specific todo by id
// @Tags todos
// @Produce json
// @Param id path int true "Todo ID"
// @Success 202 {string} string "Todo deleted"
// @Failure 404 {object} object "Todo not found"
// @Failure 500 {object} ErrorResponse "failed to delete todo"
// @Router /todos/{id} [delete]
func (app *App) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	error := ErrorResponse{"failed to delete todo"}

	deleted, err := app.db.DeleteTodoDb(id)
	if err != nil {
		log.Println("error:", err)
		c.JSON(http.StatusInternalServerError, error)
		return
	}

	if !deleted {
		log.Println("failed to delete todo because it is not found in the database")
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}

// UpdateTodo updates the todo item using its id
// @Summary Updates a todo
// @Description Changes the complete status or the title of the todo
// @Tags todos
// @Produce json
// @Param todo body Todo true "Todo body to be updated"
// @Success 201 {string} string "Todo updated successfully"
// @Failure 404 {object} object "invalid ID" or "Todo is not found"
// @Failure 400 {object} object "failed to update todo due to bad request"
// @Failure 500 {object} ErrorResponse "failed to update todo"
// @Router /todos/{id} [put]
func (app *App) UpdateTodo(c *gin.Context) {

	var todo Todo
	error := ErrorResponse{"failed to update todo"}

	// set the ID field of the todo variable explicitly
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println("error:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
		return
	}

	todo.ID = id

	if err := c.ShouldBindJSON(&todo); err != nil {
		log.Println("error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to update todo due to bad request"})
		return
	}

	updated, err := app.db.UpdateTodoDb(todo)
	if err != nil {
		log.Println("error:", err)
		c.JSON(http.StatusInternalServerError, error)
		return
	}

	if !updated {
		log.Println("failed to update todo because it is not found in the database")
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo is not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
}
