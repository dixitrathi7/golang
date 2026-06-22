package main

import (
    "fmt"
    "io"
    "net/http"
)

// Handler for GET requests to "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
    // Check if it's a GET request (optional, but good practice)
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Set response headers (e.g., content type)
    w.Header().Set("Content-Type", "text/plain")

    // Write status code (200 OK is default, but explicit here)
    w.WriteHeader(http.StatusOK)

    // Write response body
    fmt.Fprint(w, "Hello, World! This is a basic HTTP response from Go.")
}

// Handler for POST requests to "/echo" (demonstrates reading request body)
func echoHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Read the request body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close() // Always close the body to free resources

    // Set response headers
    w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)

    // Echo back the body
    fmt.Fprintf(w, "You sent: %s", string(body))
}

func main() {
    // Register handlers for specific paths
    http.HandleFunc("/", homeHandler)      // Handles GET to "/"
    http.HandleFunc("/echo", echoHandler)  // Handles POST to "/echo"

    // Start the server on port 8080
    fmt.Println("Server starting on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Server failed to start: %v\n", err)
    }
}
    // Check if it's a GET request (optional, but good practice)
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Set response headers (e.g., content type)
    w.Header().Set("Content-Type", "text/plain")

    // Write status code (200 OK is default, but explicit here)
    w.WriteHeader(http.StatusOK)

    // Write response body
    fmt.Fprint(w, "Hello, World! This is a basic HTTP response from Go.")
}

// Handler for POST requests to "/echo" (demonstrates reading request body)
func echoHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Read the request body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close() // Always close the body to free resources

    // Set response headers
    w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)

    // Echo back the body
    fmt.Fprintf(w, "You sent: %s", string(body))
}

func main() {
    // Register handlers for specific paths
    http.HandleFunc("/", homeHandler)      // Handles GET to "/"
    http.HandleFunc("/echo", echoHandler)  // Handles POST to "/echo"

    // Start the server on port 8080
    fmt.Println("Server starting on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Server failed to start: %v\n", err)
    }
}