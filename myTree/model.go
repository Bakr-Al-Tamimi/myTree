package myTree

import (
	"gopkg.in/mgo.v2/bson"
)

// Person represense collection of My Genetic Tree

type Person struct {
	ID     bson.ObjectId `bson:"_id"`
	fname  string        `json:"fname"`
	lname  string        `json:"lname"`
	bday   int32         `json:"bday"`
	bmonth int32         `json:"bmonth"`
	byear  int32         `json:"byear"`
}

// People is an array of Person
type People []Person
