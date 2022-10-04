package main

import "fmt"

type User struct {
	name     string
	age      int64
	password string
}

func change(u *User) {
	u.name = "Kate"
	u.age = 28
	u.password = "New"
}

func main() {
	//var user User = User{name: "John", age: 23, password: "pass"}
	user := User{"John", 23, "pass"}
	user.age = 45
	change(&user)
	fmt.Println(user.name, user.age)
	fmt.Println(&user)
}
