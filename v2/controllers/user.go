package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/dimitrisCBR/go-rest-tutorial/database"
	"github.com/dimitrisCBR/go-rest-tutorial/model"
	"log"
)

type (
	// UserController represents the controller for operating on the User resource
	UserController struct {
		database *mgo.Database
	}
)

// NewUserController provides a reference to a UserController with provided mongo session
func NewUserController(s *mgo.Database) *UserController {
	return &UserController{s}
}

// GetUser retrieves an individual user resource
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")
	log.Print(id)

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		log.Print(bson.IsObjectIdHex(id))
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub user
	u := model.User{}

	// Fetch user
	if err := uc.database.C(database.COL_USERS).FindId(oid).One(&u); err != nil {
		log.Print(uc.database.C(database.COL_USERS).FindId(oid).One(&u))
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}


// GetUser retrieves an individual user resource
func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request) {

	// Stub user
	u := model.Users{}

	// Fetch user
	if err := uc.database.C(database.COL_USERS).Find(nil).All(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// CreateUser creates a new user resource
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Stub an user to be populated from the body
	u := model.User{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)

	// Add an Id
	u.ID = bson.NewObjectId()

	// Write the user to mongo
	uc.database.C(database.COL_USERS).Insert(u)

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

// RemoveUser removes an existing user resource
func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request) {
	// Grab id
	id := r.FormValue("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err := uc.database.C(database.COL_USERS).RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Write status
	w.WriteHeader(200)
}
