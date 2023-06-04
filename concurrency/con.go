package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go doWork(done)
	time.Sleep(time.Second * 2)
	close(done) // cancelled child goroutine
	fmt.Println("after close done")
	time.Sleep(time.Second * 5)
}

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing infinite work")
		}
	}

}
