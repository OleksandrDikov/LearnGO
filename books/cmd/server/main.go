package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(hello))
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello World")
	if err != nil {
		log.Fatal(err)
	}
}
