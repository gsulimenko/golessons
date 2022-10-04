package main

import "fmt"

func test(someFunc func(int) int) {
	fmt.Println(someFunc(25))
}

func main() {
	square := func(x int) int {
		return x * x
	}

	test(square)
}
