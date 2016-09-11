package main

import (
	// Standard library packages
	"net/http"

	// Third party packages
	"gopkg.in/mgo.v2"
	"github.com/dimitrisCBR/shameboardAPI/v2/handlers"
	"github.com/gorilla/mux"
	"fmt"
	"io/ioutil"
	"log"
	"encoding/json"
)

type MongoConfig struct {
	Address string
	Port string
	Dbname string
}

type SigningConfig struct {
	SigningKey string
}
var signKey []byte

var config MongoConfig

func main() {

	LoadMongoConfig("./mongodatabase-conf.json")
	LoadSecurityConfig("./security-conf.json")

	fmt.Println(config)
	// Instantiate a new router
	mux := mux.NewRouter()

	mux.Handle("/get-token", handlers.GetToken(getDatabase(), signKey)).Methods("GET")

	// Get all users resource
	mux.Handle("/allshames", handlers.AuthMiddleWare(signKey).Handler(handlers.GetShames(getDatabase()))).Methods("GET")

	//Get shame by ID
	mux.Handle("/shame/{id}", handlers.AuthMiddleWare(signKey).Handler(handlers.Shame(getDatabase()))).Methods("GET")

	//Create Shame
	mux.Handle("/shame/generate", handlers.AuthMiddleWare(signKey).Handler(handlers.CreateShame(getDatabase()))).Methods("POST")

	// Get all users resource
	mux.Handle("/allusers", handlers.AuthMiddleWare(signKey).Handler(handlers.GetUsers(getDatabase()))).Methods("GET")

	// Get a user resource
	mux.Handle("/user/{id}", handlers.AuthMiddleWare(signKey).Handler(handlers.User(getDatabase()))).Methods("GET")

	// Create a new user
	mux.Handle("/user/generate", handlers.AuthMiddleWare(signKey).Handler(handlers.CreateUser(getDatabase()))).Methods("POST")

	// Remove an existing user
	mux.Handle("/user/{id}", handlers.AuthMiddleWare(signKey).Handler(handlers.DeleteUser(getDatabase()))).Methods("DELETE")

	// Fire up the server
	http.ListenAndServe(":8888", mux)
}

func LoadMongoConfig(path string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
		fmt.Println("Config File Missing. ", err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
		fmt.Println("Config Parse Error: ", err)
	}
}

func LoadSecurityConfig(path string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
		fmt.Println("Config File Missing. ", err)
	}

	var signConfig SigningConfig
	err = json.Unmarshal(file, &signConfig)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
		fmt.Println("Config Parse Error: ", err)
	}

	signKey = []byte(signConfig.SigningKey)
}

// getSession creates a new mongo session and panics if connection error occurs
func getDatabase() *mgo.Database {
	// Connect to our local mongo
	fmt.Println("mongodb://"+config.Address+":"+config.Port)
	s, err := mgo.Dial("mongodb://"+config.Address+":"+config.Port)

	// Check if connection error, is mongo running?
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Deliver session
	return s.DB(config.Dbname)
}
