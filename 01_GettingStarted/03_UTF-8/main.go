package main

import "fmt"

func main() {
	fmt.Println("\nIntger\tBinary\tHexa\tUTF-8\n")
	for i := 60; i < 123; i++ {
		fmt.Printf("%d\t%b\t%#x\t%q\n", i, i, i, i)
	}
}
