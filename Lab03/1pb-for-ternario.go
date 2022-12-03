package main

import "fmt"

func main() {
    /*
        Scrivere un programma in Go che dato un intero n, stampi tutti gli interi
        minori di n (da n a 1)
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := n; i >= 1; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
