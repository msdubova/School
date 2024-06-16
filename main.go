package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/note", getNote)

	http.ListenAndServe(":8080", mux)
}

func getNote(w http.ResponseWriter, r *http.Request) {}
