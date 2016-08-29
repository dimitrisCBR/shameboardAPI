package main

import (
	"log"
	"net/http"

	"github.com/dimitrisCBR/shameboardAPI/v2/routes"
	"gopkg.in/mgo.v2"
)

var mongoURL = "mongodb://localhost:27033"

func main() {
	router := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":8888", router))
}

func getSession() *mgo.Session {
	s, err := mgo.Dial(mongoURL)
	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

