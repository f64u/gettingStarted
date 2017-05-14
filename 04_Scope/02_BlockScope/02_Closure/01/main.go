package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "The cridets Belongs To who is in this Ring."
		fmt.Println(y)
	}

	//fmt.Println(y) // Outside The Ring
}
