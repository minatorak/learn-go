package main

import (
	"fmt"
)

func main() {
	start := true
	for start {
		start = !test()
	}
}

func test() bool {
	finit := false
	myChannel := make(chan string, 1)
	anotherChannel := make(chan string, 1)

	go func() {
		myChannel <- "cat"
	}()
	go func() {
		anotherChannel <- "cow"
	}()

	select {
	case msgFromMyChannel := <-myChannel:
		finit = true
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}
	return finit
}
