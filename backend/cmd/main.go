package main

import (
	"log"

	server "github.com/codescalersinternships/ToDoApp-Rodina/internal"
)

func main() {

	db, err := server.ConnectToDB("todoDB.db")

	if err != nil {
		log.Fatal(err)
	}

	app := server.NewApp(db)

	err = app.AppRouter()

	if err != nil {
		log.Fatal(err)
	}

}
