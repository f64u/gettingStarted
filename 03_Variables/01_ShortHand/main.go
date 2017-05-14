package main

import "fmt"

func main()  {
	
	a := "Hello World!"
	b := 42
	c := true
	d := 42.9
	e := `Do you like my Hat?`
	f := 'M'

	fmt.Printf("%v - %T \n", a, a)
	fmt.Printf("%v - %T \n", b, b)
	fmt.Printf("%v - %T \n", c, c)
	fmt.Printf("%v - %T \n", d, d)
	fmt.Printf("%v - %T \n", e, e)
	fmt.Printf("%v - %T \n", f, f)
	
}
