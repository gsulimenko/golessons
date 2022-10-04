package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("text.txt")

	if err != nil {
		fmt.Println("ОШИБКА - ", err)
		os.Exit(1)
	}

	defer file.Close()

	data := "Text to file"

	file.WriteString(data)

	file_data, _ := os.ReadFile(file.Name())

	fmt.Println(string(file_data))

	os.Remove("text.txt")

	files, err := os.ReadDir(".")

	for _, file := range files {
		fmt.Println(file.Name())
	}

}
