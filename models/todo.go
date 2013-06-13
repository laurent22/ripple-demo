package models

import (
	"database/sql"	
)

// A simple todo model
type Todo struct {
	Id int
	UserId string
	Text string
}

// Gets the tasks of the given user
func TodosByUserId(db *sql.DB, userId string) ([]*Todo, error) {
	rows, err := db.Query("SELECT id FROM todos WHERE user_id = ?", userId)
	if err != nil { return nil, err } 
	var output []*Todo
	for rows.Next() {
		var todoId int
		rows.Scan(&todoId)
		todo, _ := TodoById(db, todoId)
		output = append(output, todo)
	}
	return output, nil
}

// Gets the given todo
func TodoById(db *sql.DB, id int) (*Todo, error) {
	row := db.QueryRow("SELECT * FROM todos WHERE id = ?", id)
	
	var dbId int
	var dbUser_id string
	var dbText string
	err := row.Scan(&dbId, &dbUser_id, &dbText)
	if err != nil { return nil, err }
	
	todo := new(Todo)
	todo.Id = dbId
	todo.UserId = dbUser_id
	todo.Text = dbText
	
	return todo, nil
}

// Saves the todo to the database.
func (this *Todo) Save(db *sql.DB) error {
	if this.Id > 0 {
		db.Exec("UPDATE todos SET text = ? WHERE id = ?", this.Text, this.Id)
	} else {
		result, err := db.Exec("INSERT INTO todos(user_id, text) VALUES(?, ?)", this.UserId, this.Text)
		if err != nil { return err }
		todoId, _ := result.LastInsertId()
		this.Id = int(todoId)
	}
	
	return nil
}

// Delete the todo from the database.
func (this *Todo) Delete(db *sql.DB) error {
	db.Exec("DELETE FROM todos WHERE id = ?", this.Id)
	this.Id = 0
	return nil
}