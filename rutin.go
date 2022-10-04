package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	ch2 := make(chan int)

	go say("Hello Go", ch, ch2)
	time.Sleep(2 * time.Second)
	fmt.Println(<-ch, <-ch2)
}

func say(greet string, ch chan string, ch2 chan int) {
	fmt.Println(greet)
	ch <- "Hello from Gorutin!"
	ch2 <- 77
}
