package main

import "fmt"

func someFunc(num string) {
	fmt.Println(num)
}
func main() {
	someFunc("1")

	fmt.Println("hi")
}

// result
// 1
// hi
