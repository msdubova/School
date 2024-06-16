package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/note", getNote)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("Error happened", err.Error())
		return
	}
}

func getNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome back to school on 8080")
}
