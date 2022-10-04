package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Text to file")

	e := ioutil.WriteFile("text.txt", data, 0600)

	if e != nil {
		fmt.Println("Не могу создать файл\n", e)
	}

	file_data, err := ioutil.ReadFile("text.txt")

	if err != nil {
		fmt.Println("Не могу прочитать файл\n", err)
	}

	fmt.Println(string(file_data))

	os.Remove("text.txt")
}
