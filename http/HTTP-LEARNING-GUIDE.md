# HTTP Requests & Responses in Go - Complete Learning Guide

## 📚 Table of Contents
1. [HTTP Concepts](#http-concepts)
2. [HTTP Methods](#http-methods)
3. [Request/Response Structure](#requestresponse-structure)
4. [Status Codes](#status-codes)
5. [Go Implementation Guide](#go-implementation-guide)

---

## HTTP Concepts

### What is HTTP?
- **HTTP** = HyperText Transfer Protocol
- A **request-response protocol** used for web communication
- **Stateless**: Each request is independent (no memory of previous requests)
- **Client-Server**: Browser/App (client) requests, Web server responds

### How it Works
```
Client (Your Browser/App)
    |
    | (sends HTTP Request)
    |
    ↓
Server (Web Server)
    |
    | (processes request)
    |
    ↓ (sends HTTP Response)
    |
Client (receives response)
```

---

## HTTP Methods

### GET - Retrieve Data
- **Purpose**: Fetch data from server
- **Body**: Usually no body
- **Safe**: Yes (doesn't change data)
- **Example**: Fetching a web page, getting user info

```
GET /users?id=1 HTTP/1.1
Host: example.com
```

### POST - Send Data / Create
- **Purpose**: Send data to server, usually to create new resource
- **Body**: Contains data (form, JSON, etc.)
- **Safe**: No (changes server state)
- **Example**: Creating a new user, submitting a form

```
POST /users HTTP/1.1
Host: example.com
Content-Type: application/json

{"name": "John", "email": "john@example.com"}
```

### PUT - Update Entire Resource
- **Purpose**: Replace entire resource with new data
- **Body**: Complete new resource
- **Example**: Update entire user profile

```
PUT /users/1 HTTP/1.1
Content-Type: application/json

{"name": "Jane", "email": "jane@example.com"}
```

### DELETE - Remove Data
- **Purpose**: Delete a resource
- **Body**: Usually no body
- **Example**: Delete a user

```
DELETE /users/1 HTTP/1.1
```

### PATCH - Partial Update
- **Purpose**: Update part of a resource
- **Body**: Only fields to change
- **Example**: Update just user's email

```
PATCH /users/1 HTTP/1.1
Content-Type: application/json

{"email": "newemail@example.com"}
```

---

## Request/Response Structure

### HTTP Request Structure
```
REQUEST LINE:    GET /api/users?page=1 HTTP/1.1
HEADERS:         Host: example.com
                 Content-Type: application/json
                 Authorization: Bearer token123
                 
BODY:            (optional - for POST, PUT, PATCH)
                 {"name": "John", "age": 30}
```

### HTTP Response Structure
```
STATUS LINE:     HTTP/1.1 200 OK
HEADERS:         Content-Type: application/json
                 Content-Length: 245
                 Cache-Control: no-cache
                 
BODY:            {"id": 1, "name": "John", "age": 30}
```

### Key Request Parts
| Part | What | Example |
|------|------|---------|
| Method | Action type | GET, POST, PUT, DELETE |
| URL | Where to send | /api/users |
| Query Params | Filter data | ?id=1&name=john |
| Headers | Metadata | Content-Type, Authorization |
| Body | Data being sent | JSON, form data |

### Key Response Parts
| Part | What | Example |
|------|------|---------|
| Status Code | Success/error indicator | 200, 404, 500 |
| Headers | Response metadata | Content-Type, Set-Cookie |
| Body | Actual data returned | HTML, JSON, files |

---

## Status Codes

### 2xx - Success
- **200 OK**: Request successful, data returned
- **201 Created**: New resource created successfully
- **204 No Content**: Success but no body to return

### 3xx - Redirection
- **301 Moved Permanently**: Resource moved to new URL
- **302 Found**: Temporary redirect
- **304 Not Modified**: Use cached version

### 4xx - Client Error
- **400 Bad Request**: Invalid request format
- **401 Unauthorized**: Authentication required
- **403 Forbidden**: Authenticated but no permission
- **404 Not Found**: Resource doesn't exist
- **429 Too Many Requests**: Rate limited

### 5xx - Server Error
- **500 Internal Server Error**: Server error occurred
- **502 Bad Gateway**: Server received invalid response
- **503 Service Unavailable**: Server temporarily down

---

## Go Implementation Guide

### Key Packages
- `net/http` - Main HTTP package
- `io` - Reading response bodies
- `encoding/json` - Working with JSON
- `net/url` - URL parsing and query parameters

### Creating an HTTP Server

```go
package main

import (
    "fmt"
    "net/http"
)

// Handler function signature
func handler(w http.ResponseWriter, r *http.Request) {
    // w = ResponseWriter (where you send response)
    // r = Request (contains request data)
}

func main() {
    // Register handler for path
    http.HandleFunc("/path", handler)
    
    // Start server on port
    http.ListenAndServe(":8080", nil)
}
```

### Making HTTP Requests (Client)

```go
// Simple GET
response, err := http.Get("https://example.com/api/users")

// GET with headers
client := &http.Client{}
req, _ := http.NewRequest("GET", "https://example.com", nil)
req.Header.Set("Authorization", "Bearer token")
response, err := client.Do(req)

// POST with JSON
jsonData := `{"name": "John"}`
response, err := http.Post("https://example.com/api/users",
    "application/json", 
    bytes.NewBufferString(jsonData))
```

### Handling Requests in Server

```go
// Check request method
if r.Method == "GET" {
    // Handle GET
}

// Get query parameters
id := r.URL.Query().Get("id")

// Parse form/POST data
r.ParseForm()
name := r.FormValue("name")

// Read request body (for JSON)
var data MyStruct
json.NewDecoder(r.Body).Decode(&data)
```

### Sending Responses

```go
// Set response header
w.Header().Set("Content-Type", "application/json")

// Send response body
fmt.Fprintf(w, "Hello World")

// Set status code (must be before writing body!)
w.WriteHeader(http.StatusCreated) // 201

// Send JSON
json.NewEncoder(w).Encode(myData)

// Send error
http.Error(w, "Not found", http.StatusNotFound)
```

---

## Example Files in This Folder

1. **01-http-server.go** - Basic server (you already have this)
2. **03-http-request-methods.go** - Handling GET, POST, PUT, DELETE
3. **04-http-client-get.go** - Making GET requests to other servers
4. **05-http-client-post.go** - Making POST requests with data
5. **06-http-server-response.go** - Different response types (JSON, text, status codes)
6. **07-complete-example.go** - Full API server example

---

## Quick Start

### Run the Server
```bash
go run 01-http-server.go
```
Then visit: http://localhost:8080

### Test with curl
```bash
# GET request
curl http://localhost:8080/api/data

# POST request
curl -X POST -d "name=john" http://localhost:8080/api/data

# With custom headers
curl -H "Authorization: Bearer token" http://localhost:8080/api/data
```

---

## Key Takeaways

✅ **HTTP** is request-response based  
✅ **Methods** (GET, POST, PUT, DELETE) define the action  
✅ **Status codes** tell you if request succeeded  
✅ **Headers** carry metadata  
✅ **Body** contains the actual data  
✅ Go's `net/http` makes it easy to build servers and clients  
✅ Always **close response bodies** with `defer response.Body.Close()`  
✅ Always **set Content-Type** header in responses  

---

## Next Steps

1. Run each example file
2. Modify them and see what changes
3. Test with `curl` command
4. Create your own simple API

Good luck! 🚀
