package main

import "fmt"

const (
	_ = iota 	// 0, GONE!!!
	a = iota * 10	// 1 * 10 = 10
	b = iota * 10	// 2 * 20 = 20
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	// fmt.Println(_)
}
