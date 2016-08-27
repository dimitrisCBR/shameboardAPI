package handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/dimitrisCBR/shameboardAPI/v2/model"
	"github.com/dimitrisCBR/shameboardAPI/v2/data"
	"io/ioutil"
	"io"
)

func Index (w http.ResponseWriter, r* http.Request){
	fmt.Fprintln(w, "Welcome")
}

func Shames(w http.ResponseWriter, r *http.Request) {
	shames := model.Shames{
		model.Shame{Name: "bs"},
		model.Shame{Name: "fff"},
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
	if err := json.NewEncoder(w).Encode(data.Shames); err != nil {
		panic(err)
	}
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

	echoedShame := data.RepoCreateShame(shame)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(echoedShame); err != nil {
		panic(err)
	}
}
