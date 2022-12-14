package main

import (
	"fmt"
	"os"
)

func main() {
	x := os.Args
	fmt.Println()
	for i, k := range x {
		fmt.Println(i, k)
	}
}
