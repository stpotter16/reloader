package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
    // Parse commands and options
    if (len(os.Args) != 2) {
        log.Fatalf("Incorrect number of arguments received. Expected 1, recieved %d", len(os.Args) - 1)
    }
    filePath := os.Args[1]

    // Initialize the background go function for watching file updates

    // Initialize and run the server
    mux := http.NewServeMux()
    mux.HandleFunc("GET /", indexGet(filePath))
    http.ListenAndServe(":8080", mux)
}

func indexGet(filePath string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        htmlContent, err := readFile(filePath)
        if err != nil {
            http.Error(w, "Unable to read html file", http.StatusInternalServerError)
        }
        fmt.Fprint(w, htmlContent)
    }
}

func readFile(filePath string) (string, error) {
    data, err := os.ReadFile(filePath)
    if err != nil {
        return "", err
    }
    return string(data), nil
}
