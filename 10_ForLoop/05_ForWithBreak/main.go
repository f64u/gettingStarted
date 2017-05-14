package main

import "fmt"

func main() {

	i := 0

	/// Will Print To 9.

	for {
		fmt.Println(i)
		i++
		if i >= 10 {
			break
		}
	}
}
