package main

import "fmt"

const (
	_  = iota             // 0, GONE, Also
	KB = 1 << (iota * 10) // iota is 1
	MB = 1 << (iota * 10) // iota is 2
	GB = 1 << (iota * 10) // ....
	TB = 1 << (iota * 10)
)

func main() {
	fmt.Println("   \tBinary\t\t\t\tDecimal\n")
	fmt.Printf("KB:\t%b\t%d\n", KB, KB)
	fmt.Printf("MB:\t%b\t%d\n", MB, MB)
	fmt.Printf("GB:\t%b\t%d\n", GB, GB)
	fmt.Printf("TB:\t%b\t%d\n", TB, TB)
}
