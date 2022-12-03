package main

import "fmt"

func main() {
	/*
	   Scrivere un programma in Go che stabilisca se un intero inserito sia
       pari o dispari
	*/

	var n int
	fmt.Print("numero: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println(n, "è pari")
	} else {
		fmt.Println(n, "è dispari")
	}
}
