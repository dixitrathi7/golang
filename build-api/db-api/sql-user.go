package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB // main GORM object that represents your database connection and provides methods like: Create, Find, Update, Delete, etc.
var err error

type User struct {
	gorm.Model // gorm.Model is a struct including fields ID, CreatedAt, UpdatedAt, DeletedAt, which will be automatically managed by GORM
	//  Make sure the database schema matches my Go struct.
	Name  string `json:"name"`
	Email string `json:"email"`
}

func initDB() { // init is the spicial function in golang, it will be called before main function, so we can use it to init the database connection
	// init only run once, so we can use it to init the database connection
	dsn := "root:admin123@tcp(20.40.50.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err = db.AutoMigrate(&User{}).Error; err != nil { // autoMigrate will create the table if it does not exist, and will add missing columns, and will not delete/change current data
		// it looks at the struct and generates/updates the database table.
		log.Fatalf("failed to migrate database: %v", err)
	}
}

func userlist(w http.ResponseWriter, r *http.Request) {
	var users []User // empty slice of User struct type, to hold the result of the query

	if err := db.Find(&users).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{ // interface is store any type of value like id date and actually data
		"message": "Welcome to the UserList! of ORM DB API",
		"users":   users,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Endpoint Hit: userlist")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil { // json payload store into r.body, we decode it into user struct, if error, return 400
		// decode read the json payload from the request body and decode it into the user struct. If there is an error during decoding, it returns a 400 Bad Request error.
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

	fmt.Println("Endpoint Hit: createUser")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := db.Delete(&User{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	fmt.Println("Endpoint Hit: deleteUser")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var input User // read the request body and decode it into a user struct.
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user User // find the user by id, if not found, return 404
	if err := db.First(&user, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	user.Name = input.Name
	user.Email = input.Email // only go struct is updated at this point, we need to save it back to the database

	if err := db.Save(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	fmt.Println("Endpoint Hit: updateUser")
}
