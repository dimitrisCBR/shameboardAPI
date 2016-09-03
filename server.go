package main

import (
	// Standard library packages
	"net/http"

	// Third party packages
	"gopkg.in/mgo.v2"
	"github.com/dimitrisCBR/shameboardAPI/v2/database"
	"github.com/dimitrisCBR/shameboardAPI/v2/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Instantiate a new router
	mux := mux.NewRouter()

	mux.Handle("/get-token", handlers.GetToken(getDatabase())).Methods("GET")

	// Get all users resource
	mux.Handle("/allshames", handlers.JwtMiddleware.Handler(handlers.GetShames(getDatabase()))).Methods("GET")

	//Get shame by ID
	mux.Handle("/shame/{id}", handlers.JwtMiddleware.Handler(handlers.Shame(getDatabase()))).Methods("GET")

	//Create Shame
	mux.Handle("/shame/generate", handlers.JwtMiddleware.Handler(handlers.CreateShame(getDatabase()))).Methods("POST")

	// Get all users resource
	mux.Handle("/allusers", handlers.JwtMiddleware.Handler(handlers.GetUsers(getDatabase()))).Methods("GET")

	// Get a user resource
	mux.Handle("/user/{id}", handlers.JwtMiddleware.Handler(handlers.User(getDatabase()))).Methods("GET")

	// Create a new user
	mux.Handle("/user", handlers.JwtMiddleware.Handler(handlers.CreateUser(getDatabase()))).Methods("POST")

	// Remove an existing user
	mux.Handle("/user/{id}", handlers.JwtMiddleware.Handler(handlers.DeleteUser(getDatabase()))).Methods("DELETE")

	// Fire up the server
	http.ListenAndServe("localhost:8888", mux)
}

// getSession creates a new mongo session and panics if connection error occurs
func getDatabase() *mgo.Database {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost:27033")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s.DB(database.DB_NAME)
}
