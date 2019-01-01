package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Code    string `json:"code,omitempty"`
	Credits int    `json:"credits,omitempty"`
}

var courses []Course

func getAllCourses(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(courses)
}

func main() {
	router := mux.NewRouter()

	courses = append(courses, Course{ID: "1", Name: "Cálculo diferencial", Code: "CB-101", Credits: 4})
	courses = append(courses, Course{ID: "2", Name: "Geometría analítica", Code: "CB-201", Credits: 2})

	router.HandleFunc("/api/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/api/courses/{id}", getAllCourses).Methods("GET")
	if err := http.ListenAndServe(":5000", router); err != nil {
		log.Fatal(err)
	}
}
