package main

import "fmt"

type User struct {
	name     string
	age      int64
	password string
}

func (u User) getName() string {
	return u.name
}

func (u *User) setName(name1 string) {
	u.name = name1
}

func (u User) isElder() bool {
	a := u.age
	isTrue := false

	if a >= 18 {
		isTrue = true
	} else if a < 18 {
		isTrue = false
	}

	return isTrue
}

func main() {
	user := User{"John", 25, "pass"}
	fmt.Println(user.getName())
	user.setName("Kate")
	fmt.Println(user.getName())
	fmt.Println(user.isElder())
}
