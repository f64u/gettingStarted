package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {

		case i % 12 == 0:
			fmt.Printf("%d -- FizzBuzz\n", i)
		case i % 3 == 0:
			fmt.Printf("%d -- Fizz\n", i)
		case i % 4 == 0:
			fmt.Printf("%d -- Buzz\n", i)
		default:
			println(i)
		}
	}
}
