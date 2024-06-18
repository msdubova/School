package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Student struct {
	Name      string
	Id        int
	UpdatedAt time.Time
}

type Class []Student

var John = Student{
	Name: "John",
	Id:   2323,
}

var Emma = Student{
	Name: "Emma",
	Id:   3232,
}

var David = Student{
	Name: "David",
	Id:   6676,
}

var toddlers = Class{John, Emma, David}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /class", checkAuth(processClass))
	mux.HandleFunc("GET /student/{id}", checkAuth(processStudent))
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("Error happened", err.Error())
		return
	}

}

func processClass(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getClass(w, r)
	}
}
func processStudent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getStudent(w, r)
	}
}
func getClass(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(toddlers)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	searchingId := r.PathValue("id")
	id, err := strconv.Atoi(searchingId)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, student := range toddlers {
		if student.Id == id {
			err := json.NewEncoder(w).Encode(student)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
	}
	fmt.Println("студент не знайден")
	w.WriteHeader(http.StatusNotFound)
}

type User struct {
	Name string
	Role string
}

var teacher = User{
	Name: "MrBond",
	Role: "teacher",
}

func checkAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name, role, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if teacher.Name != name || teacher.Role != role {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
