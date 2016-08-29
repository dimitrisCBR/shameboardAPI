package handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/dimitrisCBR/shameboardAPI/v2/model"
	"io/ioutil"
	"io"
	"log"
	"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
	"github.com/dimitrisCBR/shameboardAPI/v2/database"
)

func Index (w http.ResponseWriter, r* http.Request){
	fmt.Fprintln(w, "Welcome")
}

func Shames(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r,"database").(*mgo.Session)

	// load the shames
	var shames []*model.Shame
	if err := db.DB(database.DB_NAME).C(database.COL_SHAMES).
		Find(nil).Sort("-when").Limit(100).All(&shames); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(shames); err != nil {
		panic(err)
	}

}

func Shame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shameId := vars["shame_id"]
	fmt.Fprintln(w, "Shame:", shameId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//if err := json.NewEncoder(w).Encode(data.Shames); err != nil {
	//	panic(err)
	//}
}

func ShameCreate(w http.ResponseWriter, r *http.Request){
	var shame model.Shame
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &shame); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

}

func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				WriteError(w, ErrInternalServer)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

