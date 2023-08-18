# ToDo App

This is a todo app, which is a web-based user interface, where users may interact with their tasks and manage their todo list. The frontend is designed with Vue.js and interfaces with the backend API designed in Go.

# Table of Contents
1. [App Features](#app-features)
2. [Used Technologies](#used-technologies)
3. [Installation](#installation)
4. [Manual](#manual)
5. [API Documentation](#api-documentation)
7. [Testing](#testing)

   
## App Features

1. Add new todos to help users organize their tasks using the add button.
2. Mark the completed todos as done using a checkbox.
3. Delete unwanted todos using the delete button.
4. Rename a todo by double-clicking it and then typing the new name.
   
## Used Technologies

 ### Backend

- Golang: A quick and efficient backend programming language.
- Gin: A Go-based lightweight web framework used to develop the API.
- Swagger: An API documentation and testing tool.

 ### Frontend

- Vue.js: A modern framework for creating user interfaces in JavaScript or Typescript. In this app, Typescript is used.
- Cypress: An end-to-end testing framework for web applications.

## Installation
Before starting, ensure that the necessary tools are installed:

- [Go](https://go.dev/)
- [Node.js](https://nodejs.org/en)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Cypress](https://www.cypress.io/)

## Manual
- Clone the repository:
```sh
$ git clone https://github.com/codescalersinternships/ToDoApp-Rodina.git
```
- Go to the repository directory:
```sh
$ cd ToDoApp-Rodina
```
- Backend Setup:
```sh
$ cd backend

$ go mod download

$ go run main.go

```
- Frontend Setup:
```sh
$ cd frontend

$ npm install

$ npm run serve
```
### __How to run with docker?__

Run the app using docker as follows:
```sh
$ docker-compose up -d
```
### __How to access the app?__
- Open your web browser and go to http://localhost:8080 to access the app.

## API Documentation
After starting the backend server, navigate to http://localhost:8096/swagger/index.html to view the API documentation and test the backend API using Swagger.

## Testing
After installing all dependencies, you can test either the frontend or the backend as follows:

 ### Backend
```sh
$ cd backend
$ go mod download
$ go test ./...
```
### Frontend
```sh
$ cd frontend
$ npm install
$ npm run test
```

