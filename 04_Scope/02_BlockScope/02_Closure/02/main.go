package main

import "fmt"

func increment() int {
	x++
	return x
}

var x = 0

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}
