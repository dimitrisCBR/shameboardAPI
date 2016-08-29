package model

import  "gopkg.in/mgo.v2/bson"

type Shame struct {
	ID bson.ObjectId `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Source User `json:"source"`
	Destination User `json:"destination"`
}

type Shames []Shame