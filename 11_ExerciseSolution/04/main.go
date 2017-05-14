package main

import "fmt"

func main() {
	var xnum, lnum int

	fmt.Print("Enter a large number: ")
	fmt.Scan(&xnum)

	fmt.Print("Enter a smaller number: ")
	fmt.Scan(&lnum)

	fmt.Println("The remainder by dividing them equals", xnum % lnum)
}
