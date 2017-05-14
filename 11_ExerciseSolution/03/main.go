package main

import (
	"fmt"
	"bufio"
	"os"
)




func main() {

	in := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Your name: ")


	name, _ := in.ReadString('\n')
	name = name[:len(name) - 1]

	fmt.Printf("Hello %s!", name)
}
