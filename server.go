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


	// Get all users resource
	mux.HandleFuncC(pat.Get("/allshames"), handlers.GetShames(getDatabase()))

	//Get shame by ID
	mux.HandleFuncC(pat.Get("/shame/:id"), handlers.Shame(getDatabase()))

	//Create Shame
	mux.HandleFuncC(pat.Post("/shame/generate"), handlers.ShameCreate(getDatabase()))

	//// Get all users resource
	//r.Handle("/allusers", handlers.GetShames(getDatabase())).Methods("GET")
	//
	//// Get a user resource
	//r.Handle("/user/:id", uc.GetUser).Methods("GET")
	//
	//// Create a new user
	//r.Handle("/user", uc.CreateUser).Methods("POST")
	//
	//// Remove an existing user
	//r.Handle("/user/:id", uc.RemoveUser).Methods("DELETE")
	//
	//// Get a ShameController instance
	//sc := controllers.NewShameController(getDatabase())
	//
	//// Get a shame resource
	//r.Handle("/shame/:id", sc.GetShame).Methods("GET")
	//
	//// Create a new shame
	//r.Handle("/shame", sc.CreateShame).Methods("POST")
	//
	//// Remove an existing shame
	//r.Handle("/shame/:id", sc.RemoveShame).Methods("DELETE")

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
