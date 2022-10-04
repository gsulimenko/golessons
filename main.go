package main

import (
	"fmt"
)

func main() {
	/*	names := [3]string{"John", "Jordan", "Kate"}
		fmt.Println(names, len(names))

		for i := 0; i < len(names); i++ {
			fmt.Println(names[i])
		}*/

	slice := []int{3, 1, 2, 5, 7, 4}
	fmt.Println(slice)
	slice = append(slice, 0)
	//sort.Ints(slice)
	fmt.Println(slice)

	for _, el := range slice {
		//fmt.Printf("%d: %d\n", i, el)
		fmt.Println(el)
	}

}
