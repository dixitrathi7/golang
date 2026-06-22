package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// This example shows how to MAKE HTTP GET requests (as a client)
// Unlike the server examples, this is how to request data FROM another server

func main() {
	// Example 1: Simple GET request to a public API
	fmt.Println("=== Example 1: Simple GET Request ===")

	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Make the GET request
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error making request:", err)
	}
	defer response.Body.Close() // Important: Always close the response body

	// Check the status code
	fmt.Printf("Status Code: %d\n", response.StatusCode)
	fmt.Printf("Status: %s\n\n", response.Status)

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}

	fmt.Println("Response Body:")
	fmt.Println(string(body))

	// Example 2: GET request with query parameters
	fmt.Println("\n=== Example 2: GET with Query Parameters ===")

	queryURL := "https://jsonplaceholder.typicode.com/posts?userId=1&_limit=2"

	resp2, err := http.Get(queryURL)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer resp2.Body.Close()

	body2, _ := io.ReadAll(resp2.Body)
	fmt.Println("Posts for user 1 (limited to 2):")
	fmt.Println(string(body2))

	// Example 3: GET request with custom headers
	fmt.Println("\n=== Example 3: GET with Custom Headers ===")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://httpbin.org/headers", nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	// Add custom headers
	req.Header.Add("User-Agent", "MyGoApp/1.0")
	req.Header.Add("Authorization", "Bearer mytoken123")

	resp3, err := client.Do(req)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer resp3.Body.Close()

	body3, _ := io.ReadAll(resp3.Body)
	fmt.Println("Response with custom headers:")
	fmt.Println(string(body3))
}
