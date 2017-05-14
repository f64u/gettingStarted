package main

import "fmt"

func main() {
	for i := 0 ; i <= 256; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
	}

	foo := "a"
	fmt.Println(foo)
	fmt.Printf("%T\n", foo)
}
