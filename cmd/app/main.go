package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type options struct {
    filePath string
    port int
}

func main() {
    // Parse commands and options
    options := parseOptions()

    // Initialize the background go function for watching file updates

    // Initialize and run the server
    mux := http.NewServeMux()
    mux.HandleFunc("GET /", indexGet(options.filePath))
    addr := fmt.Sprintf(":%d", options.port)
    log.Printf("Serving on %s", addr)
    http.ListenAndServe(addr, mux)
}

func parseOptions() options {
    // TODO: Error handling
    file := flag.String("file", "", "The file to serve")
    port := flag.Int("port", 8080, "The port to serve the file from on localhost")
    flag.Parse()
    return options{
        filePath: *file,
        port: *port,
    }
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
        log.Printf("Error reading file: %v", err)
        return "", err
    }
    return string(data), nil
}
