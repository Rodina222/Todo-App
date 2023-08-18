# Todo App Frontend

This is the frontend of the todo app, which is a web-based user interface, where users may interact with their tasks and manage their todo list. This frontend is designed with Vue.js and interfaces with the backend API.

## Features

1. Add new todos to help users organize their tasks using the add button.
2. Mark the completed todos as done using a checkbox.
3. Delete unwanted todos using the delete button.
4. Rename a todo by double-clicking it and then typing the new name.

# User Manual
- Clone the repository:
```sh
$ git clone https://github.com/codescalersinternships/ToDoApp-Rodina.git
```
- Go to the repository directory:
```sh
$ cd ToDoApp-Rodina/frontend
```
- Project setup:
```
npm install
```
- Compiles and hot-reloads for development:
```
npm run serve
```
- Compiles and minifies for production:
```
npm run build
```
- Lints and fixes files:
```
npm run lint
```

### __How to test?__
- Run components tests as follows: 
```
npx cypress run --component
```
- Run integration tests as follows:
```
npx cypress run or npx cypress open
```
If all tests pass, the output should indicate that the tests were successful; if any test fails, the output will indicate which test failed.

