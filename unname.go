package main

import "fmt"

func change(str *string) {
	*str = "LOL"
}

func main() {
	s := "LLL"
	fmt.Println(s)
	change(&s)
	fmt.Println(s)
}
