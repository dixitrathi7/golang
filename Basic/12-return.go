package main

import "fmt"

func checkusername(username string) string {
	if len(username) < 5 {
		return "Username is too short"
	} else if len(username) > 15 {
		return "Username is too long"
	} else {
		return "Username is valid"
	}
}

func main() {

	username := "dixit"
	check := checkusername(username)
	fmt.Println("Your Username : ", username, "status : ", check)

}
