package main

import "fmt"

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			fmt.Println("sliceToChannel ", n)
			out <- n
		}
		close(out) // wait channel
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			fmt.Println("sq ", n)
			out <- n * n
		}
		close(out) // wait channel
	}()
	return out
}

func main() {

	// input
	nums := []int{2, 3, 4, 5, 7, 9}
	dataChannel1 := sliceToChannel(nums)

	finalChannel1 := sq(dataChannel1)

	for n := range finalChannel1 {
		fmt.Println(n)
	}

}
