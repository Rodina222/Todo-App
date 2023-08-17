package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodo(t *testing.T) {

	tempDir := t.TempDir()

	dbPath := filepath.Join(tempDir, "todoDB.db")

	database, err := ConnectToDB(dbPath)
	assert.NoError(t, err)

	app := NewApp(database)

	router := gin.Default()

	router.POST("/todos", app.CreateTodo)

	t.Run("create todo successfully", func(t *testing.T) {

		newTodo := Todo{
			Title:     "Todo",
			Completed: false,
		}
		requestBody, err := json.Marshal(newTodo)
		assert.NoError(t, err)

		request, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(requestBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusCreated, recorder.Code, "got %d status code but want status code 201", recorder.Code)

	})

	t.Run("failed to create todo due to bad request", func(t *testing.T) {

		requestBody := []byte(`{"id":1,"title":"todo","completed": }`)

		req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(requestBody))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "got %d status code but want status code 400", rec.Code)

	})

}

func TestGetAllTodos(t *testing.T) {

	tempDir := t.TempDir()

	dbPath := filepath.Join(tempDir, "todoDB.db")

	db, err := ConnectToDB(dbPath)
	assert.NoError(t, err)

	app := NewApp(db)

	router := gin.Default()

	router.GET("/todos", app.GetAllTodos)

	req, err := http.NewRequest("GET", "/todos", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "got %d status code but want status code 200", recorder.Code)

}

func TestDeleteTodo(t *testing.T) {

	tempDir := t.TempDir()

	dbPath := filepath.Join(tempDir, "todoDB.db")

	database, err := ConnectToDB(dbPath)
	assert.NoError(t, err)

	app := NewApp(database)

	router := gin.Default()

	router.DELETE("/todos/:id", app.DeleteTodo)

	t.Run("delete todo successfully", func(t *testing.T) {

		newTodo := Todo{
			ID:        15,
			Title:     "Todo2",
			Completed: false,
		}

		db := &DB{db: database}

		id, err := db.CreateTodo(newTodo)
		assert.NoError(t, err)
		fmt.Println("id", id)

		req, err := http.NewRequest("DELETE", "/todos/"+strconv.FormatInt(id, 10), nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "got %d status code but want status code 200", rec.Code)

	})

	t.Run("failed to delete todo due to invalid id", func(t *testing.T) {

		req, err := http.NewRequest("DELETE", "/todos/5", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code, "got %d status code but want status code 404", rec.Code)

	})

}

func TestUpdateTodo(t *testing.T) {

	tempDir := t.TempDir()

	dbPath := filepath.Join(tempDir, "todoDB.db")

	database, err := ConnectToDB(dbPath)
	assert.NoError(t, err)

	app := NewApp(database)

	router := gin.Default()

	router.PUT("/todos/:id", app.UpdateTodo)

	t.Run("update todo successfully", func(t *testing.T) {

		newTodo := Todo{
			ID:        1,
			Title:     "Todo5",
			Completed: false,
		}

		db := &DB{db: database}

		id, err := db.CreateTodo(newTodo)
		assert.NoError(t, err)

		fmt.Println("id", id)

		updatedTodo := Todo{
			ID:        1,
			Title:     "Todo5",
			Completed: true,
		}

		requestBody, err := json.Marshal(updatedTodo)
		assert.NoError(t, err)

		req, err := http.NewRequest("PUT", "/todos/"+strconv.FormatInt(id, 10), bytes.NewBuffer(requestBody))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "got %d status code but want status code 200", rec.Code)

	})

	t.Run("failed to update todo due to invalid id", func(t *testing.T) {

		updatedTodo := Todo{
			ID:        100,
			Title:     "Todo5",
			Completed: true,
		}

		requestBody, err := json.Marshal(updatedTodo)
		assert.NoError(t, err)

		req, err := http.NewRequest("PUT", "/todos/"+strconv.Itoa(updatedTodo.ID), bytes.NewBuffer(requestBody))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code, "got %d status code but want status code 404", rec.Code)

	})

	t.Run("failed to update todo due to bad request", func(t *testing.T) {

		requestBody := []byte(`{"id":1,"title":"todo","completed": }`)

		req, err := http.NewRequest("PUT", "/todos/1", bytes.NewBuffer(requestBody))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "got %d status code but want status code 400", rec.Code)

	})
}

func TestGetTodoByID(t *testing.T) {

	tempDir := t.TempDir()

	dbPath := filepath.Join(tempDir, "todoDB.db")

	database, err := ConnectToDB(dbPath)
	assert.NoError(t, err)

	app := NewApp(database)
	assert.NoError(t, err)

	router := gin.Default()

	router.GET("/todos/:id", app.GetTodoByID)

	t.Run("get todo successfully", func(t *testing.T) {

		newTodo := Todo{
			ID:        1,
			Title:     "Todo2",
			Completed: false,
		}

		db := &DB{db: database}

		id, err := db.CreateTodo(newTodo)
		assert.NoError(t, err)
		fmt.Println("id", id)

		req, err := http.NewRequest("GET", "/todos/"+strconv.Itoa(newTodo.ID), nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code, "got %d status code but want status code 200", recorder.Code)

	})

	t.Run("failed to get todo due to invalid id", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/todos/50", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusNotFound, recorder.Code, "got %d status code but want status code 404", recorder.Code)

	})

}
