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


func GetShames (db *mgo.Database) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {

		// Stub shame Array
		shames := model.Shames{}

		// Fetch shames
		if err := db.C(database.COL_SHAMES).Find(nil).All(&shames); err != nil {
			w.WriteHeader(404)
			ErrorWithJSON(w, "ID Not found", http.StatusBadRequest)
			return
		}

		// Marshal provided interface into JSON structure
		uj, _ := json.Marshal(shames)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(uj)
	}
}


func Shame (db *mgo.Database) http.HandlerFunc {
	return func ( w http.ResponseWriter, r *http.Request) {
		// Grab id
		var params = mux.Vars(r)
		shameId := params["id"]

		// Verify id is ObjectId, otherwise bail
		if !bson.IsObjectIdHex(shameId) {
			w.WriteHeader(404)
			ErrorWithJSON(w, "ID Not found", http.StatusBadRequest)
			return
		}

		// Grab id
		oid := bson.ObjectIdHex(shameId)

		// Stub shame
		shame := model.Shame{}

		// Fetch user
		if err := db.C(database.COL_SHAMES).FindId(oid).One(&shame); err != nil {
			w.WriteHeader(404)
			return
		}

		// Marshal provided interface into JSON structure
		returnedShames, _ := json.Marshal(shame)

		// Write content-type, statuscode, payload
		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(returnedShames)

	}
}


func CreateShame (db *mgo.Database) http.HandlerFunc {
	return func ( w http.ResponseWriter, r *http.Request) {

		// Stub a shame to be populated from the body
		shame := model.Shame{}

		// Populate the shame data
		err := json.NewDecoder(r.Body).Decode(&shame)
		if(err != nil) {
			ErrorWithJSON(w, "Error in the provided object", http.StatusBadRequest)
			return
		}

		// Add an Id
		shame.ID = bson.NewObjectId()

		// Write the user to mongo
		if err := db.C(database.COL_SHAMES).Insert(shame); err != nil {
			if mgo.IsDup(err) {
				ErrorWithJSON(w, "duplicate object", http.StatusBadRequest)
				return
			}

			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			return
		}

		// Marshal provided interface into JSON structure
		returnedShames, _ := json.Marshal(shame)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		w.Write(returnedShames)
	}
}

