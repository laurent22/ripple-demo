package rippledemo

import (
	"../models"
	"../ripple"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

// A simple controller to GET, POST, PUT and DELETE todos.
// The database connection is injected in the controller constructor.
type TodoController struct {
	db *sql.DB
}

func NewTodoController(db *sql.DB) *TodoController {
	output := new(TodoController)
	output.db = db
	return output
}

func (this *TodoController) Get(ctx *ripple.Context) {
	// Get the todo ID
	todoId, _ := strconv.Atoi(ctx.Params["id"])
	// Get the model
	todo, _ := models.TodoById(this.db, todoId)
	// Return it
	ctx.Response.Body = todo
}

func (this *TodoController) Post(ctx *ripple.Context) {
	// Get the todo JSON
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	// Unserialize it
	var todo models.Todo
	json.Unmarshal(body, &todo)
	
	// Save the new todo
	todo.Save(this.db)
	// Return what we've just saved
	ctx.Response.Body = todo
}

func (this *TodoController) Put(ctx *ripple.Context) {
	// Get the todo JSON
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	// Unserialize it
	var todo models.Todo
	json.Unmarshal(body, &todo)

	// Get the todo we want to update
	todoId, _ := strconv.Atoi(ctx.Params["id"])
	todo.Id = todoId
	
	// Update the model
	todo.Save(this.db)
	
	// Return the updated model
	ctx.Response.Body = todo
}

func (this *TodoController) Delete(ctx *ripple.Context) {
	// Get the todo ID
	todoId, _ := strconv.Atoi(ctx.Params["id"])
	// Get the model
	todo, _ := models.TodoById(this.db, todoId)
	// Delete it
	todo.Delete(this.db)
}