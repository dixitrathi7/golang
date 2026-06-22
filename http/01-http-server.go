package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprint(w, "This is GET")
		return
	}

	if r.Method == "POST" {
		fmt.Fprint(w, "This is POST")
		return
	}

	fmt.Fprintf(w, "Method not allowed")
}

func loginhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprint(w, "This is GETon login")
		return
	}

	if r.Method == "POST" {
		fmt.Fprint(w, "This is POST on login")
		return
	}
	fmt.Fprintf(w, "Method not allowed on login")
}

func main() {
	http.HandleFunc("/path", handler)

	http.HandleFunc("/login", loginhandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
