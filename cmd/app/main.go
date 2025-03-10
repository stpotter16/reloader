package main

import (
	"fmt"
	"net/http"
)

func main() {
    // Parse commands and options
    // Initialize the background go function for watching file updates
    // Initialize and run the server
    mux := http.NewServeMux()
    mux.HandleFunc("GET /", indexGet())
    http.ListenAndServe(":8080", mux)
}

func indexGet() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello, World!")
    }
}
