package dao

import (
	"log"

	. "github.com/heberqc/horariojs-backend/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CoursesDAO sets the connection parameters
type CoursesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// The name of the collecion
const (
	COLLECTION = "courses"
)

// Connect function stablish connection with the database
func (m *CoursesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// FindAll returns all the courses
func (m *CoursesDAO) FindAll() ([]Course, error) {
	var courses []Course
	err := db.C(COLLECTION).Find(bson.M{}).All(&courses)
	return courses, err
}

// FindByID returns an specific course
func (m *CoursesDAO) FindByID(id string) (Course, error) {
	var course Course
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&course)
	return course, err
}
