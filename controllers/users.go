package rippledemo

import (
	"../models"
	"../ripple"
	"database/sql"
)

type UserController struct {
	db *sql.DB
}

func NewUserController(db *sql.DB) *UserController {
	output := new(UserController)
	output.db = db
	return output
}

func (this *UserController) GetTodos(ctx *ripple.Context) {
	userId, _ := ctx.Params["id"]
	todos, _ := models.TodosByUserId(this.db, userId)
	ctx.Response.Body = todos
}