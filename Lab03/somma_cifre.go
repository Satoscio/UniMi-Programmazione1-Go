/*
	Scrivere un programma somma_cifre.go che calcola la somma
	delle cifre di un numero intero positivo fornito da standard input.
*/

package main

import "fmt"

func main() {
	var n, somma int
	fmt.Scan(&n)
	for {
		somma += n % 10
		n /= 10
		if n == 0 {
			break
		}
	}
	fmt.Println("Somma cifre:", somma)
}
