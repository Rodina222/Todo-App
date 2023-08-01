package main

import (
	"errors"
	"flag"
	"log"

	server "github.com/codescalersinternships/ToDoApp-Rodina/backend/internal"
	"github.com/gin-gonic/gin"
)

// ErrDBPathNotFound is returned if the user doesn't mention the path of the database
var ErrDBPathNotFound = errors.New("database path must be provided after the -db flag")

func main() {

	gin.SetMode(gin.ReleaseMode)

	var dbPath string

	flag.StringVar(&dbPath, "db", "./todoapp.db", "mention the filepath of the database")

	flag.Parse()

	if dbPath == "" {
		log.Fatal(ErrDBPathNotFound)
	}

	db, err := server.ConnectToDB(dbPath)

	if err != nil {
		log.Fatal(err)
	}

	app, err := server.NewApp(db)

	if err != nil {
		log.Fatal(err)
	}

	if err = app.Run(); err != nil {

		log.Fatal(err)
	}

	defer db.Close()

}
