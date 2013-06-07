package main

import (
	"database/sql"
	"log"
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
	_, err := initializeDatabase()
	if err != nil {
		log.Panicln(err)
	}
}




// // This is a simple REST API demo that runs on localhost, port 8080. To start it:
// //
// //     go run demo.go
// //
// // Then in a different terminal window:
// //
// //     go run demo_client.go
// //
// // You can also check the URLs directly in a browser. For instance:
// //
// //     http://localhost:8080/users/1

// package main

// import (
// 	"../ripple"
// 	"./controllers"
// 	"net/http"
// )

// func main() {
// 	// Build the REST application

// 	app := ripple.NewApplication()

// 	// Create a controller and register it. Any number of controllers
// 	// can be registered that way.

// 	userController := rippledemo.NewUserController()
// 	app.RegisterController("users", userController)

// 	// Setup the routes. The special patterns `_controller` will automatically match
// 	// an existing controller, as defined above. Likewise, `_action` will match any
// 	// existing action.

// 	app.AddRoute(ripple.Route{Pattern: ":_controller/:id/:_action"})
// 	app.AddRoute(ripple.Route{Pattern: ":_controller/:id/"})
// 	app.AddRoute(ripple.Route{Pattern: ":_controller"})

// 	// Start the server

// 	http.ListenAndServe(":8080", app)
// }
