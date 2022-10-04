package main

import (
	"fmt"
)

func main() {
	var money map[string]int = map[string]int{
		"dollars": 1000,
		"euros":   2000,
		"rubels":  3000,
	}

	el, status := money["euros"]

	fmt.Println(el, status)
}
