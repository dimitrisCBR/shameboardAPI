package main

import (
	// Standard library packages
	"net/http"

	// Third party packages
	"gopkg.in/mgo.v2"
	"goji.io"
	"goji.io/pat"
	"github.com/dimitrisCBR/shameboardAPI/v2/database"
	"github.com/dimitrisCBR/shameboardAPI/v2/handlers"
)

func main() {
	// Instantiate a new router
	mux := goji.NewMux()

	mux.Handle(pat.Get("/get-token"), handlers.GetToken(getDatabase()))

	// Get all users resource
	mux.Handle(pat.Get("/allshames"), handlers.JwtMiddleware.Handler(handlers.GetShames(getDatabase())))

	//Get shame by ID
	mux.Handle(pat.Get("/shame/:id"), handlers.JwtMiddleware.Handler(handlers.Shame(getDatabase())))

	//Create Shame
	mux.Handle(pat.Post("/shame/generate"), handlers.JwtMiddleware.Handler(handlers.CreateShame(getDatabase())))

	// Get all users resource
	mux.Handle(pat.Get("/allusers"), handlers.JwtMiddleware.Handler(handlers.GetUsers(getDatabase())))

	// Get a user resource
	mux.Handle(pat.Get("/user/:id"), handlers.JwtMiddleware.Handler(handlers.User(getDatabase())))

	// Create a new user
	mux.Handle(pat.Post("/user"), handlers.JwtMiddleware.Handler(handlers.CreateUser(getDatabase())))

	// Remove an existing user
	mux.Handle(pat.Delete("/user/:id"), handlers.JwtMiddleware.Handler(handlers.DeleteUser(getDatabase())))

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
