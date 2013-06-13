package main

import (
	"./controllers"
	"./ripple"
	"database/sql"
	"log"
	"net/http"
	_ "./go-sqlite3"
	"io/ioutil"
	"strings"
	"os"
)

func initializeDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./Todos.sqlite3")
	if err != nil { return nil, err }
	
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY AUTOINCREMENT, user_id TEXT, text TEXT)")
	if err != nil { return nil, err }
	
	return db, nil
}

// Very simple HTTP handler for testing the web application and API
func webAppHandler(w http.ResponseWriter, r *http.Request) {
	t := strings.Split(r.URL.Path, "/")
	filePath := "html/" + t[len(t) - 1]
	
	fi, err := os.Stat(filePath)
	if fi == nil && err != nil {
		// File doesn't exist
		return
	}
	
	content, _ := ioutil.ReadFile(filePath)
	w.Write(content)
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
	
	// Handle the front-end
	http.HandleFunc("/app/", webAppHandler)
	
	// Handle the REST API	
	app.SetBaseUrl("/api/")
	http.HandleFunc("/api/", app.ServeHTTP)
	
	log.Println("Starting server...")
	http.ListenAndServe(":8080", nil)
}