package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http" // browser sends an HTTP request: receives it and parses it for you.

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func onlyPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: onlyPost")
	fmt.Fprintf(w, "This endpoint only accepts POST requests")

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/home", homePage)
	router.HandleFunc("/art", allArticles).Methods("GET")
	router.HandleFunc("/post", onlyPost).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router)) // Start listening on port 8080 and wait for requests
}

func main() {
	handleRequests()
}
