package model

import  "gopkg.in/mgo.v2/bson"

type Shame struct {
	ID bson.ObjectId `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Source User `json:"source" bson:"source"`
	Destination User `json:"destination" bson:"destination"`
}

type Shames []Shame

type ShameCollection struct {
	Data []Shame `json:"data"`
}

type ShameResource struct {
	Data Shame `json:"data"`
}

var USER_ID = "id"
var USER_NAME = "name"
var USER_EMAIL = "email"
var USER_PASSWORD = "password"

type User struct {
	ID bson.ObjectId  `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Users []User

type UserCollection struct {
	Data []User `json:"data"`
}

type UsereResource struct {
	Data User `json:"data"`
}
