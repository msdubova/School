package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/note", getNote)
}

func getNote(w http.ResponseWriter, r *http.Request) {}
