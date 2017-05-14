package main

import "fmt"

func main() {
	var sum int
	for i := 0; i < 100000000000; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}

	fmt.Println("Sum is:", sum)
}
