package main

import (
	"fmt"
)

func main() {
	myChannel := make(chan string) // make channel and data flow in channel is string

	go func() { // gorountine
		myChannel <- "dataTest" // send "dataTest" to channel
	}()

	msg := <-myChannel // blocking line of code @main
	// for read data from channel close or recived data from channel

	fmt.Println(msg)

}
