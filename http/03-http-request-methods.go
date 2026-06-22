package main

import (
	"fmt"
	"log"
	"net/http"
)

// This example shows how to handle different HTTP request methods

func requestHandlerv1(w http.ResponseWriter, r *http.Request) {
	// r.Method tells you which HTTP method was used
	switch r.Method {
	case "GET":
		// Client is asking for data
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "You sent a GET request to: %s\n", r.URL.Path)
		fmt.Fprintf(w, "Query parameters: %s\n", r.URL.RawQuery)

	case "POST":
		// Client is sending data to create something
		w.Header().Set("Content-Type", "text/plain")

		// Parse the form data from request body
		r.ParseForm()

		fmt.Fprintf(w, "You sent a POST request\n")
		fmt.Fprintf(w, "Form data: %v\n", r.Form)

	case "PUT":
		// Client is updating data
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "You sent a PUT request (for updating data)\n")

	case "DELETE":
		// Client is requesting to delete data
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "You sent a DELETE request (for deleting data)\n")

	default:
		// Unsupported method
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Create a simple server that handles different request types
	http.HandleFunc("/api/data", requestHandlerv1)

	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("\nTry these in another terminal:")
	fmt.Println("  GET:    curl http://localhost:8080/api/data?name=john&age=25")
	fmt.Println("  POST:   curl -X POST -d \"name=john&age=25\" http://localhost:8080/api/data")
	fmt.Println("  PUT:    curl -X PUT http://localhost:8080/api/data")
	fmt.Println("  DELETE: curl -X DELETE http://localhost:8080/api/data")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
