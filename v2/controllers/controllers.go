package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/dimitrisCBR/shameboardAPI/v2/model"
	"github.com/dimitrisCBR/shameboardAPI/v2/database"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var DBName = "test";

type (
	// UserController represents the controller for operating on the User resource
	ShameController struct{
		session *mgo.Session
	}
)

func NewUserController(s *mgo.Session) *ShameController {
	return &ShameController{s}
}

// GetUser retrieves an individual user resource
func (sc ShameController) GetShame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shameId := vars["id"]

	// Grab id
	oid := bson.ObjectIdHex(shameId)

	// Stub object
	shame := model.Shame{}

	sc.session.DB(database.DB_NAME).C(database.COL_SHAMES).FindId(oid).One(&shame);

	// Marshal provided interface into JSON structure
	shameResponse, _ := json.Marshal(shame)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", shameResponse)
}

// CreateUser creates a new Shame resource
func (sc ShameController) CreateShame(w http.ResponseWriter, r *http.Request) {
	// Stub to be populated from the body
	shame := model.Shame{}

	// Populate the data
	json.NewDecoder(r.Body).Decode(&shame)

	shame.ID = -bson.NewObjectId()

	sc.session.DB(database.DB_NAME).C(database.COL_SHAMES).Insert(shame)

	// Marshal provided interface into JSON structure
	shameResponse, _ := json.Marshal(shame)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", shameResponse)
}

// RemoveUser removes an existing user resource
func (sc ShameController) RemoveUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r);
	id := vars["id"]

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err := sc.session.DB(database.DB_NAME).C(database.COL_USERS).RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Write status
	w.WriteHeader(200)
}
