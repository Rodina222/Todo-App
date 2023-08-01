package internal

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectToDB(t *testing.T) {

	_, err := sql.Open("sqlite3", "database.db")
	assert.NoError(t, err)

}

func TestCreateTable(t *testing.T) {

	database, err := sql.Open("sqlite3", "database.db")
	assert.NoError(t, err)

	db := &DB{db: database}

	err = db.CreateTable()
	assert.NoError(t, err)

}
