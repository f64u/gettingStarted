package main

import (
	"fmt"
	"github.com/fady-adel/gettingStarted/04_Scope/01_PackageScope/02_Visibility/vis"
)

func main() {
	fmt.Println(vis.MyName + "\n")
	vis.PrintVar()
}
