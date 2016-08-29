package main

import (
	// Standard library packages
	"net/http"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	"github.com/dimitrisCBR/go-rest-tutorial/controllers"
	"gopkg.in/mgo.v2"
	"github.com/dimitrisCBR/go-rest-tutorial/database"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Get a UserController instance
	uc := controllers.NewUserController(getDatabase())

	// Get all users resource
	r.GET("/allusers", uc.GetUsers)

	// Get a user resource
	r.GET("/user/:id", uc.GetUser)

	// Create a new user
	r.POST("/user", uc.CreateUser)

	// Remove an existing user
	r.DELETE("/user/:id", uc.RemoveUser)

	// Get a ShameController instance
	sc := controllers.NewShameController(getDatabase())

	// Get a shame resource
	r.GET("/shame/:id", sc.GetShame)

	// Create a new shame
	r.POST("/shame", sc.CreateShame)

	// Remove an existing shame
	r.DELETE("/shame/:id", sc.RemoveShame)


	// Fire up the server
	http.ListenAndServe("localhost:8888", r)
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
