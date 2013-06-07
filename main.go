package main

import (
	"./controllers"
	"./ripple"
	"database/sql"
	"log"
	"net/http"
	_ "./go-sqlite3"
)

func initializeDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./Todos.sqlite3")
	if err != nil { return nil, err }
	
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY AUTOINCREMENT, user_id TEXT, text TEXT)")
	if err != nil { return nil, err }
	
	return db, nil
}

func main() {
	db, err := initializeDatabase()
	if err != nil {
		log.Panicln(err)
	}
	defer db.Close()
	
	app := ripple.NewApplication()
	
	todoController := rippledemo.NewTodoController(db)
	userController := rippledemo.NewUserController(db)

	app.RegisterController("todos", todoController)
	app.RegisterController("users", userController)
	
	app.AddRoute(ripple.Route{Pattern: ":_controller/:id/:_action"})
	app.AddRoute(ripple.Route{Pattern: ":_controller/:id/"})
	app.AddRoute(ripple.Route{Pattern: ":_controller"})
	
	http.ListenAndServe(":8080", app)
}