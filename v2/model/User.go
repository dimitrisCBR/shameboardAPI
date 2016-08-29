package model

import  "gopkg.in/mgo.v2/bson"

var USER_ID = "id"
var USER_NAME = "name"
var USER_EMAIL = "email"
var USER_PASSWORD = "password"

type User struct {
	ID bson.ObjectId  `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Users []User
