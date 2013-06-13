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
	// Get the user ID
	userId, _ := ctx.Params["id"]
	// Get the todos of this user
	todos, _ := models.TodosByUserId(this.db, userId)
	// Return them
	ctx.Response.Body = todos
}