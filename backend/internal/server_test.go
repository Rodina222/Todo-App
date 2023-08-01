package internal

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodo(t *testing.T) {

	db, err := ConnectToDB("./todoDB.db")
	assert.NoError(t, err)

	app, err := NewApp(db)
	assert.NoError(t, err)

	router := gin.Default()

	router.POST("/todos", app.CreateTodo)

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

}

func TestGetAllTodos(t *testing.T) {

	db, err := ConnectToDB("./todoDB.db")
	assert.NoError(t, err)

	app, err := NewApp(db)
	assert.NoError(t, err)

	router := gin.Default()

	router.GET("/todos", app.GetAllTodos)

	req, err := http.NewRequest("GET", "/todos", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "got %d status code but want status code 200", recorder.Code)

}

func TestDeleteTodo(t *testing.T) {

	db, err := ConnectToDB("./todoDB.db")
	assert.NoError(t, err)

	app, err := NewApp(db)
	assert.NoError(t, err)

	router := gin.Default()

	router.DELETE("/todos/:id", app.DeleteTodo)

	id := 15

	req, err := http.NewRequest("DELETE", "/todos/"+strconv.Itoa(id), nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusAccepted, recorder.Code, "got %d status code but want status code 202", recorder.Code)

}

func TestUpdateTodo(t *testing.T) {

	db, err := ConnectToDB("./todoDB.db")
	assert.NoError(t, err)

	app, err := NewApp(db)
	assert.NoError(t, err)

	router := gin.Default()

	router.PUT("/todos/:id", app.UpdateTodo)

	updatedTodo := Todo{
		ID:        5,
		Title:     "Todo3",
		Completed: true,
	}

	requestBody, err := json.Marshal(updatedTodo)
	assert.NoError(t, err)

	req, err := http.NewRequest("PUT", "/todos/"+strconv.Itoa(updatedTodo.ID), bytes.NewBuffer(requestBody))
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "got %d status code but want status code 200", recorder.Code)
}

func TestGetTodoByID(t *testing.T) {

	db, err := ConnectToDB("./todoDB.db")
	assert.NoError(t, err)

	app, err := NewApp(db)
	assert.NoError(t, err)

	router := gin.Default()

	router.POST("/todos", app.CreateTodo)

	newTodo := Todo{
		ID:        12,
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

	router.GET("/todos/:id", app.GetTodoByID)

	req, err := http.NewRequest("GET", "/todos/"+strconv.Itoa(newTodo.ID), nil)
	assert.NoError(t, err)

	recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

}
