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

	mux.HandleFunc("/class", checkAuth(processNote))

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
		// case http.MethodPut:
		// 	putNote(w, r)
	}
}

func getNote(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(toddlers)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//	func putNote(w http.ResponseWriter, r *http.Request) {
//		err := json.NewDecoder(r.Body).Decode(&John)
//		if err != nil {
//			fmt.Println(err.Error())
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//		John.UpdatedAt = time.Now()
//		err = json.NewEncoder(w).Encode(John)
//		if err != nil {
//			fmt.Println(err.Error())
//			w.WriteHeader(http.StatusInternalServerError)
//			return
//		}
//	}
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
