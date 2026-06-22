package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// This example shows how to BUILD proper HTTP responses on the server side

// User struct for JSON responses
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Example 1: Simple text response
func textHandler(w http.ResponseWriter, r *http.Request) {
	// Set response header (tells client what type of data we're sending)
	w.Header().Set("Content-Type", "text/plain")

	// Write response body
	fmt.Fprintf(w, "Hello! This is a plain text response.\n")
	fmt.Fprintf(w, "You requested: %s\n", r.URL.Path)
}

// Example 2: JSON response
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	// Create data
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	}

	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Encode struct to JSON and write to response
	json.NewEncoder(w).Encode(user)
}

// Example 3: JSON array response
func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "John", Email: "john@example.com"},
		{ID: 2, Name: "Jane", Email: "jane@example.com"},
		{ID: 3, Name: "Bob", Email: "bob@example.com"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Example 4: Response with status codes
func statusHandler(w http.ResponseWriter, r *http.Request) {
	// Different responses based on path
	if r.URL.Path == "/ok" {
		w.WriteHeader(http.StatusOK) // 200
		fmt.Fprintf(w, "Success! Status code: 200")

	} else if r.URL.Path == "/created" {
		w.WriteHeader(http.StatusCreated) // 201
		fmt.Fprintf(w, "Resource created! Status code: 201")

	} else if r.URL.Path == "/notfound" {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "Resource not found! Status code: 404")

	} else if r.URL.Path == "/error" {
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, "Server error! Status code: 500")

	} else {
		// Using http.Error for error responses
		http.Error(w, "Bad Request", http.StatusBadRequest) // 400
	}
}

// Example 5: Response with custom headers
func headerHandler(w http.ResponseWriter, r *http.Request) {
	// Set multiple headers
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Custom-Header", "MyCustomValue")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Authorization", "Bearer token123")

	fmt.Fprintf(w, "<html><body>")
	fmt.Fprintf(w, "<h1>Response with Custom Headers</h1>")
	fmt.Fprintf(w, "<p>Check the response headers in your browser</p>")
	fmt.Fprintf(w, "</body></html>")
}

// Example 6: Redirects
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect to another URL
	http.Redirect(w, r, "https://www.google.com", http.StatusMovedPermanently) // 301
	// Or use http.StatusFound (302) for temporary redirect
}

func main() {
	http.HandleFunc("/text", textHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/users", usersHandler)

	http.HandleFunc("/ok", statusHandler)
	http.HandleFunc("/created", statusHandler)
	http.HandleFunc("/notfound", statusHandler)
	http.HandleFunc("/error", statusHandler)
	http.HandleFunc("/badrequest", statusHandler)

	http.HandleFunc("/headers", headerHandler)
	http.HandleFunc("/redirect", redirectHandler)

	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("\nTry these URLs:")
	fmt.Println("  http://localhost:8080/text")
	fmt.Println("  http://localhost:8080/json")
	fmt.Println("  http://localhost:8080/users")
	fmt.Println("  http://localhost:8080/ok")
	fmt.Println("  http://localhost:8080/created")
	fmt.Println("  http://localhost:8080/notfound")
	fmt.Println("  http://localhost:8080/headers")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
