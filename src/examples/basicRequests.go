package main

import (
	"fmt"
	"net/http"
)

// Create handler
type myHandler struct {
	greeting string
}

// Implement handler interface
func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v world", mh.greeting)))
}

func main() {
	// Initialize myHandler and set greeting
	// Every request route not registered will go fall to this.
	http.Handle("/", &myHandler{greeting: "Hello"})

	// More streamlined, less flexibility though
	http.HandleFunc("/func", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Handle Func"))
	})

	http.ListenAndServe(":8080", nil) // with nil it will use DefaultServerMux
}
