package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Student struct {
	Name      string
	Id        int
	UpdatedAt time.Time
}

type Class []Student

var John = Student{
	Name:      "John",
	Id:        2323,
	UpdatedAt: time.Now(),
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/class", processNote)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error happened", err.Error())
		return
	}

}

func processNote(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getNote(w, r)
	case http.MethodPut:
		putNote(w, r)
	}
}

func getNote(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(John)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func putNote(w http.ResponseWriter, r *http.Request) {

}
