package main

import (
	"database/sql"
	"log"
	_ "./go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./Todos.sqlite3")
	if err != nil {
		log.Panicln(err)
	}
	defer db.Close()
	
	_, err = db.Exec("CREATE TABLE \"todos\" (\"id\" INTEGER PRIMARY KEY AUTOINCREMENT, \"user_id\" TEXT, \"text\" TEXT)")
	if err != nil {
		log.Panicln(err)
	}
}
