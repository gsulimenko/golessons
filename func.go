package main

import "fmt"

func main() {
	a := func(x int, y int) int {
		return x + y
	}
	sum := a(3, 4)
	fmt.Println(sum)
}
