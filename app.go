package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/heberqc/horariojs-backend/config"
	. "github.com/heberqc/horariojs-backend/dao"
)

var config = Config{}
var dao = CoursesDAO{}

func getAllCourses(w http.ResponseWriter, req *http.Request) {
	courses, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, courses)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/api/courses/{id}", getAllCourses).Methods("GET")
	if err := http.ListenAndServe(":5000", router); err != nil {
		log.Fatal(err)
	}
}
