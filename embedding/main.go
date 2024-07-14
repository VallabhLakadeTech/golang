package main

import "fmt"

type User struct {
	name  string
	email string
}

func (user *User) notify() {
	fmt.Println("Name and Email: ", user.name, user.email)
}

type Admin struct {
	*User
	level int
}

func main() {

	admin := Admin{
		User: &User{
			name:  "vallabh",
			email: "vallabh.lakade.1990@gmail.com",
		},
		level: 1,
	}
	admin.User.notify()
	admin.notify()

}
