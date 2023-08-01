# ToDo App Backend

This is a todo app that supports the following features: 

1. Assist users in organising their tasks into todos.
2. Mark the completed todos using a checkbox.
3. Remove unnecessary or unwanted todos.
4. Rename a todo.

## __API Endpoints:__
- `/GET /todos:` gets a list of all todos.
- `/GET /todos/id:` gets a specific todo by id.
- `POST /todos:` creates a new todo.
- `DELETE /todos/id:` deletes a specific todo by id. 
- `PUT /todos/id:` updates a specific todo by id.


## __Manual:__

1. Clone the repository:
```sh
$ git clone https://github.com/codescalersinternships/ToDoApp-Rodina.git 
```
2. Go to the repository directory:
```sh
$ cd ToDoApp-Rodina/backend
```
3. Install dependencies:
```sh
$ go get -d ./...
```
4. Build the package:
```sh
$ go build -o "bin/app" cmd/main.go
```
 ### __How to use without docker?__

1. Run the app as follows:
```sh
$ ./bin/app -db [dbPath]
```
Notes:
- You need to move the binary to any `$PATH` directory first.
- You must give the database path or it will be treated as an error.

2. Get all todos as follows: 
- `/GET /todos:` gets a list of all todos.
```sh
$ http://localhost:8096/todos
```
3. Get a todo by id as follows:
- `/GET /todos/id:` gets a specific todo by id.
```sh
$ http://localhost:8096/todos/id
```
4. Create a new todo as follows:
- `POST /todos:` creates a new todo.
```sh
$ http://localhost:8096/todos
```
5. Delete a todo as follows:
- `DELETE /todos/id:` deletes a specific todo by id. 
```sh
$ http://localhost:8096/todos/id
```
6. Update a todo by id as follows:
- `PUT /todos/id:` updates a specific todo by id.
```sh
$ http://localhost:8096/todos/id
```

### __How to use docker?__

Run the app using docker as follows:
```sh
$ docker-compose up -d
```

### __How to test?__

Run all the tests as follows: 
```sh
go test ./....
```
If all tests pass on, the result should show that the tests were successful as follows:
```sh
PASS
ok      github.com/codescalersinternships/backend/ToDoApp-Rodina/internal       0.091s
```
If any test fails, the output will tell you which test failed.