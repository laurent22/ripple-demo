This is a todo demo for [Ripple](https://github.com/laurent22/ripple), the REST API framework for Go.

The demo creates a SQLlite database and defines two controllers:

    controllers/users.go
    controllers/todos.go
    
and one model:

    models/todo.go    

The front end itself is a JS application that fetches, creates and deletes tasks.

# Installation

    git clone https://github.com/laurent22/ripple-demo.git
    git submodule init
    git submodule update
    
# Running the demo

    go run main.go
    
The front-end application will then be at `http://localhost/app/:8080`

And the REST API itself is at `http://localhost/api/:8080` and defines the following paths:

    GET /users/:userId/todos
    GET /todos
    GET /todos/:todoId
    POST /todos
    PUT /todos/:todoId
    DELETE /todos/:todoId