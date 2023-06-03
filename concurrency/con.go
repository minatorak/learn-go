package main

import (
	"fmt"
	"time"
)

func main() {
	unbufferedChannel := make(chan string, 0)
	chars := []string{"a", "b", "c"}

	go func() {
		for _, c := range chars {
			fmt.Println("goroutine sender ", c)
			unbufferedChannel <- c
		}
	}()

	n := len(chars)
	for n >= 1 {
		time.Sleep(time.Second * 1)
		fmt.Println("start receive")
		receive := <-unbufferedChannel
		fmt.Println("receive ", receive)
		n -= 1
	}
}
