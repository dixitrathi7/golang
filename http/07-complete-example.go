package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// Product represents a product in our store
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Our simple "database" of products
var products = []Product{
	{ID: 1, Name: "Laptop", Price: 999.99},
	{ID: 2, Name: "Mouse", Price: 29.99},
	{ID: 3, Name: "Keyboard", Price: 79.99},
}

// GET /products - List all products
func listProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// GET /product?id=1 - Get single product
func getProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get query parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	// Convert string to int
	productID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	// Find product
	for _, p := range products {
		if p.ID == productID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

// POST /product - Create new product
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate new ID
	newProduct.ID = len(products) + 1
	products = append(products, newProduct)

	// Return 201 Created status
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)
}

// Make a request to external API
func externalAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Make request to JSONPlaceholder API
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/users/1", nil)

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

// Health check endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status":"healthy","message":"Server is running"}`)
}

func main() {
	// Server routes
	http.HandleFunc("/products", listProductsHandler)
	http.HandleFunc("/product", getProductHandler)
	http.HandleFunc("/create", createProductHandler)
	http.HandleFunc("/external-user", externalAPIHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("🚀 Complete HTTP Server Example")
	fmt.Println("================================")
	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("\n📝 Try these commands:\n")

	fmt.Println("1. GET all products:")
	fmt.Println("   curl http://localhost:8080/products\n")

	fmt.Println("2. GET single product:")
	fmt.Println("   curl http://localhost:8080/product?id=1\n")

	fmt.Println("3. CREATE new product:")
	fmt.Println(`   curl -X POST http://localhost:8080/create \`)
	fmt.Println(`     -H "Content-Type: application/json" \`)
	fmt.Println(`     -d '{"name":"Monitor","price":299.99}'\n`)

	fmt.Println("4. GET external API data:")
	fmt.Println("   curl http://localhost:8080/external-user\n")

	fmt.Println("5. Health check:")
	fmt.Println("   curl http://localhost:8080/health\n")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
