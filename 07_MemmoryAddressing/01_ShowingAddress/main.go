package main

import "fmt"

func main() {
	a := 42
	fmt.Println("a -", a)
	fmt.Println("The address of a is" , &a)
	fmt.Printf("%#x in decimal equals %d\n", &a, &a)
}
