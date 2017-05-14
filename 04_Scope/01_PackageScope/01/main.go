package main

import "fmt"

var x string = "3+3"

func main() {
	fmt.Println(x)
	foo()
}

func foo()  {
	fmt.Println(x)
}
