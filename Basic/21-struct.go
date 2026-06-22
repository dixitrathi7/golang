package main

import "fmt"

type User struct {
	Username   string
	LoginCount int
}

func IncreaseLoginCount(u User) int {
	u.LoginCount++
	return u.LoginCount
}

func IncreaseLoginCountPtr(u *User) {
	u.LoginCount++
}

func main() {

	user1 := User{Username: "dixit", LoginCount: 0}

	//	fmt.Println(IncreaseLoginCount(user1))
	IncreaseLoginCountPtr(&user1)

	fmt.Println(user1.LoginCount)

}
