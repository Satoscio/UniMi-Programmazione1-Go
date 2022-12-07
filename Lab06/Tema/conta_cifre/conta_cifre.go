package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// La variabile di ritorno bool richiesta è inutile
func contaCifre(s string, numCifre *[10]int) {
	for _, r := range s {
		if unicode.IsDigit(r) {
			x, _ := strconv.Atoi(string(r))
			numCifre[x]++
		}
	}
}

func main() {
	var s string
	var cifre = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var numCifre = [10]int{}
	for {
		fmt.Scan(&s)
		if s == "stop" {
			break
		}
		contaCifre(s, &numCifre)
	}
	fmt.Println("Cifre:\t", cifre)
	fmt.Println("Trovate:", numCifre)
}
