package handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/dimitrisCBR/shameboardAPI/v2/model"
	"log"
	"gopkg.in/mgo.v2"
	"github.com/dimitrisCBR/shameboardAPI/v2/database"
	"gopkg.in/mgo.v2/bson"
	"goji.io"
	"golang.org/x/net/context"
	"goji.io/pat"
)


func GetShames (db *mgo.Database) goji.HandlerFunc {
	return func (context context.Context, w http.ResponseWriter, r *http.Request) {

		// Stub shame Array
		shames := model.Shames{}

		// Fetch shames
		if err := db.C(database.COL_SHAMES).Find(nil).All(&shames); err != nil {
			w.WriteHeader(404)
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


func Shame (db *mgo.Database) goji.HandlerFunc {
	return func (context context.Context, w http.ResponseWriter, r *http.Request) {
		// Grab id
		shameId := pat.Param(context, "id")
		fmt.Println(shameId)

		// Verify id is ObjectId, otherwise bail
		if !bson.IsObjectIdHex(shameId) {
			w.WriteHeader(404)
			ErrorWithJSON(w, "Nothing found with provided id", http.StatusBadRequest)
			return
		}

		// Grab id
		oid := bson.ObjectIdHex(shameId)

		// Stub shame
		shame := model.Shame{}

		// Fetch user
		if err := db.C(database.COL_SHAMES).FindId(oid).One(&shame); err != nil {
			w.WriteHeader(404)
			fmt.Println("not found in db")
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


func ShameCreate (db *mgo.Database) goji.HandlerFunc {
	return func (context context.Context, w http.ResponseWriter, r *http.Request) {

		// Stub shame Array
		shames := model.Shames{}

		// Fetch shames
		if err := db.C(database.COL_SHAMES).Find(nil).All(&shames); err != nil {
			w.WriteHeader(404)
			return
		}

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
				ErrorWithJSON(w, "Book with this ISBN already exists", http.StatusBadRequest)
				return
			}

			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed insert book: ", err)
			return
		}



		// Fetch shames
		if err := db.C(database.COL_SHAMES).Find(nil).All(&shames); err != nil {
			w.WriteHeader(404)
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

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}


func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
			//	WriteError(w, ErrInternalServer)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

