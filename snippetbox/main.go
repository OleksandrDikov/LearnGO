package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Snippet View"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Snippet Create"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create/{$}", snippetCreate)

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":http", mux)
	log.Fatal(err)
}
