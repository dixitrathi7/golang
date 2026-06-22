package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// This example shows how to MAKE HTTP POST requests (send data to a server)

func main() {
	// Example 1: POST with form data
	fmt.Println("=== Example 1: POST with Form Data ===")

	// Create form data
	formData := url.Values{}
	formData.Set("title", "My New Post")
	formData.Set("body", "This is the content of my post")
	formData.Set("userId", "1")

	// Make POST request with form data
	response, err := http.PostForm(
		"https://jsonplaceholder.typicode.com/posts",
		formData,
	)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	fmt.Printf("Status: %d\n", response.StatusCode)
	fmt.Println("Response:")
	fmt.Println(string(body))

	// Example 2: POST with JSON data
	fmt.Println("\n=== Example 2: POST with JSON Data ===")

	jsonData := `{
		"title": "My JSON Post",
		"body": "This is JSON content",
		"userId": 1
	}`

	// Create request with JSON body
	req, err := http.NewRequest(
		"POST",
		"https://jsonplaceholder.typicode.com/posts",
		bytes.NewBufferString(jsonData),
	)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Set Content-Type header for JSON
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer resp.Body.Close()

	body2, _ := io.ReadAll(resp.Body)
	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Println("Response:")
	fmt.Println(string(body2))

	// Example 3: POST with custom headers
	fmt.Println("\n=== Example 3: POST with Custom Headers ===")

	req3, _ := http.NewRequest(
		"POST",
		"https://httpbin.org/post",
		bytes.NewBufferString("message=hello"),
	)

	// Add multiple headers
	req3.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req3.Header.Set("Authorization", "Bearer token123")
	req3.Header.Set("X-Custom-Header", "CustomValue")

	resp3, _ := client.Do(req3)
	defer resp3.Body.Close()

	body3, _ := io.ReadAll(resp3.Body)
	fmt.Println("Response:")
	fmt.Println(string(body3))
}
