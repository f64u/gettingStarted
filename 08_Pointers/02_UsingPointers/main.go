package main

import "fmt"

func main() {

	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 45
	fmt.Println(a)
}
