package main

import "fmt"

func test(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	test("Hello!")()
}
