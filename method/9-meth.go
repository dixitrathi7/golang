package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func (u User) getinfo() string {
	return fmt.Sprintf("ID: %d, Name: %s, Email: %s, Age: %d", u.ID, u.Name, u.Email, u.Age)
}

func (u *User) updateEmail(newEmail string) {
	u.Email = newEmail
}

func (u *User) updateAge() {
	u.Age++
}

func main() {
	user := &User{ID: 1, Name: "Dixit", Email: "dixit@gmail.com", Age: 24}

	fmt.Println(user.getinfo())

	user.updateEmail("dixitrathi@gamil.com")

	fmt.Println(user.getinfo())

	user.updateAge()

	fmt.Println(user.getinfo())
}
