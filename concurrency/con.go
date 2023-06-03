package main

import (
	"fmt"
)

func main() {
	bufferedChannel := make(chan string, 1)
	chars := []string{"a", "b", "c", "d", "e"}

	go func() {
		for _, c := range chars {
			fmt.Println("goroutine sender ", c)
			bufferedChannel <- c
		}
	}()

	n := len(chars)
	for n >= 1 {
		// time.Sleep(time.Second * 1)
		fmt.Println("start receive")
		receive := <-bufferedChannel
		fmt.Println("receive ", receive)
		n -= 1
	}
}
