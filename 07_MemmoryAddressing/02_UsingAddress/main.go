package main

import "fmt"

const MetersToYards float64 = 1.09361

func main() {
	var meters float64;
	fmt.Print("Enter what you swam in meters: ")
	fmt.Scan(&meters)
	yards := meters * MetersToYards
	fmt.Println(meters, "meters is", yards, "yards.")

}
