package main

import "fmt"

func main() {
	x := 0

	fmt.Println(x)

	foo()
}

func foo()  {
	fmt.Println(x)
}
