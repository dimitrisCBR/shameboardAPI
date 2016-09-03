package handlers


import (
	"net/http"
	"encoding/json"
	"github.com/dimitrisCBR/shameboardAPI/v2/model"
	"gopkg.in/mgo.v2"
	"github.com/dimitrisCBR/shameboardAPI/v2/database"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
)


func GetUsers (db *mgo.Database) http.HandlerFunc {
	return func ( w http.ResponseWriter, r *http.Request) {

		// Stub user Array
		users := model.Users{}

		// Fetch users
		if err := db.C(database.COL_USERS).Find(nil).All(&users); err != nil {
			w.WriteHeader(404)
			ErrorWithJSON(w, "database error", http.StatusBadRequest)
			return
		}

		// Marshal provided interface into JSON structure
		uj, _ := json.Marshal(users)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(uj)
	}
}


func User (db *mgo.Database) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		// Grab id
		var params = mux.Vars(r)
		id := params["id"]

		// Verify id is ObjectId, otherwise bail
		if !bson.IsObjectIdHex(id) {
			w.WriteHeader(404)
			ErrorWithJSON(w, "Not valid id", http.StatusBadRequest)
			return
		}

		// Grab id
		oid := bson.ObjectIdHex(id)

		// Stub user
		user := model.User{}

		// Fetch user
		if err := db.C(database.COL_USERS).FindId(oid).One(&user); err != nil {
			w.WriteHeader(404)
			ErrorWithJSON(w, "Nothing found with provided id", http.StatusBadRequest)
			return
		}

		// Marshal provided interface into JSON structure
		returnedUsers, _ := json.Marshal(user)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(returnedUsers)

	}
}


func CreateUser (db *mgo.Database) http.HandlerFunc {
	return func ( w http.ResponseWriter, r *http.Request) {


		// Stub a user to be populated from the body
		user := model.User{}

		// Populate the shame data
		err := json.NewDecoder(r.Body).Decode(&user)
		if(err != nil) {
			ErrorWithJSON(w, "Error in the provided object", http.StatusBadRequest)
			return
		}

		// Add an Id
		user.ID = bson.NewObjectId()

		// Write the user to mongo
		if err := db.C(database.COL_USERS).Insert(user); err != nil {
			if mgo.IsDup(err) {
				ErrorWithJSON(w, "duplicate object", http.StatusBadRequest)
				return
			}

			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			return
		}


		returnedUser, _ := json.Marshal(user)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		w.Write(returnedUser)
	}
}



func DeleteUser (db *mgo.Database) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {

// No - op
	}
}

