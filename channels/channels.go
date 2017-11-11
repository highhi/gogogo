package main

import (
	"fmt"
)

func observer() (chan string, func(string)) {
	ch := make(chan string)
	return ch, func(msg string) { ch <- msg }
}

func main() {
	ch, ob := observer()
	go ob("hi")
	fmt.Println(<-ch)
	go ob("fo")
	fmt.Println(<-ch)
}
