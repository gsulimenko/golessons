package main

import "fmt"

func first() {
	fmt.Println("Hello from Function!")
}

func sum(a int, b int) int {
	result := a + b
	return result
}

func main() {
	first()
	x := sum(5, 7)
	fmt.Println(x)
}
