package models

import "gopkg.in/mgo.v2/bson"

// Course representation, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Course struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	Name    string        `bson:"name" json:"name"`
	Code    string        `bson:"code" json:"code"`
	Credits int           `bson:"credits" json:"credits"`
}
