package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "stringa"
	ns := strings.Split(s, " ")
	for _, p := range ns {
		fmt.Printf("%T : %s\n", p, p)
	}
}