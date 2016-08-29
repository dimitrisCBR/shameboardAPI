package database

//import (
//	"gopkg.in/mgo.v2"
//	"github.com/dimitrisCBR/shameboardAPI/v2/model"
//)

var dbURL = "http://localhost:27033"
var DB_NAME = "shameboard"
var COL_USERS = "user"
var COL_SHAMES = "shames"
var COL_BADGES = "badges"
//
//var db *mgo.Database
//
//type MongoDBConn struct {
//	session *mgo.Session
//}
//
//func NewMongoDBConn() *MongoDBConn {
//	return &MongoDBConn{}
//}
//
//func (m *MongoDBConn) Connect() *mgo.Session {
//	session, err := mgo.Dial(dbURL)
//	if err != nil {
//		panic(err)
//	}
//	m.session = session
//	return m.session
//}
//
//func (m *MongoDBConn) Stop() {
//	m.session.Close()
//}
//
//func (m *MongoDBConn) AddShame(name, description string) (err error) {
//	c := m.session.DB("test").C("shames")
//	err = c.Insert(&model.Shame{name, description})
//	if err != nil {
//		panic(err)
//	}
//	return nil
//}
//
//func (m *MongoDBConn) AddUser(name, email, password string) (err error) {
//	c := m.session.DB("test").C("shames")
//	err = c.Insert(&model.User{name, email, password})
//	if err != nil {
//		panic(err)
//	}
//	return nil
//}