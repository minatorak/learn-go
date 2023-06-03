package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}
func main() {
	go someFunc("1") // fork to child
	go someFunc("2")
	go someFunc("3")

	fmt.Println("hi")
	time.Sleep(time.Second * 2)
}
