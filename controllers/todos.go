package rippledemo

import (
	"../models"
	"../ripple"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type TodoController struct {
	db *sql.DB
}

func NewTodoController(db *sql.DB) *TodoController {
	output := new(TodoController)
	output.db = db
	return output
}

func (this *TodoController) Get(ctx *ripple.Context) {
	todoId, _ := strconv.Atoi(ctx.Params["id"])
	todo, _ := models.TodoById(this.db, todoId)
	ctx.Response.Body = todo
}

func (this *TodoController) Post(ctx *ripple.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	var todo models.Todo
	json.Unmarshal(body, &todo)
	
	todo.Save(this.db)
	ctx.Response.Body = todo
}

func (this *TodoController) Put(ctx *ripple.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	var todo models.Todo
	json.Unmarshal(body, &todo)

	todoId, _ := strconv.Atoi(ctx.Params["id"])
	todo.Id = todoId
	
	todo.Save(this.db)
	ctx.Response.Body = todo
}

func (this *TodoController) Delete(ctx *ripple.Context) {
	todoId, _ := strconv.Atoi(ctx.Params["id"])
	todo, _ := models.TodoById(this.db, todoId)
	todo.Delete(this.db)
}