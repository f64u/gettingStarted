package main

import "fmt"

const (
	a = iota // 0
	b = iota // 1
	c = iota

)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
