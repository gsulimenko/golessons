package main

import "fmt"

func main() {
	var name string
	var age int8
	fmt.Println("What's your name?")
	fmt.Scan(&name)
	fmt.Println("How old are you?")
	fmt.Scan(&age)
	fmt.Println("Hello " + name + ", you are " + fmt.Sprint(age) + " years old!")
	if age <= 25 {
		fmt.Print("Hey, dude!")
	} else if (age > 25) && (age < 50) {
		fmt.Println("Hello, sir!")
	} else {
		fmt.Println("You are to old to be here! :)")
	}

}
