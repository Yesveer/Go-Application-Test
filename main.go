package main

import (
    "fmt"
    "net/http"
)

// Handler function for the home route
func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World! Welcome to my Go app.")
}

func main() {
    // Route configuration
    http.HandleFunc("/", homeHandler)

    // Run the server on port 5001
    fmt.Println("Server is running on port 5001...")
    if err := http.ListenAndServe(":5001", nil); err != nil {
        fmt.Println("Error starting the server:", err)
    }
}
