package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage! of ORM DB API")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/home", homePage).Methods("GET")
	router.HandleFunc("/userlist", userlist).Methods("GET")
	router.HandleFunc("/deleteUser", deleteUser).Methods("DELETE")
	router.HandleFunc("/updateUser", updateUser).Methods("PUT")
	router.HandleFunc("/createUser", createUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {

	fmt.Println("Starting ORM DB API server on port 8080...")

	initDB() // Initialize the database connection and perform any necessary setup, such as creating tables or seeding initial data.
	handleRequests()
}
