package main

import "fmt"

func main() {
	a := 5
	pointer := &a
	fmt.Println(pointer)
	fmt.Println(*pointer)
}
