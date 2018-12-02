package myTree

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Repository
type Repository struct{}

// SERVER the DB server
const SERVER = "localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "myTree"

// DOCNAME the name of the document
const DOCNAME = "people"

// GetPeople returns the list of People
func (r Repository) GetPeople() People {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		log.Fatalln("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := People{}
	if err := c.Find(nil).All(&results); err != nil {
		log.Fatalln("Failed to write results:", err)
	}
	return results
}

// AddPerson insert person into DB
func (r Repository) AddPerson(person Person) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	person.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(person)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// UpdatePerson updates infor about person in DB(not used for now)
func (r Repository) UpdatePerson(person Person) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	session.DB(DBNAME).C(DOCNAME).UpdateId(person.ID, person)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// DeletePerson deletes Person from DB(not used for now)
func (r Repository) DeletePerson(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	// Varify ID is ObjectID, otherwise error
	if !bson.IsObjectIdHex(id) {
		return "ID NOT FOUND"
	}
	// Grab ID
	oid := bson.ObjectIdHex(id)
	// Remove user with ID
	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "INTERNAL ERROR"
	}
	// Write status
	return "PERSON REMOVED"
}