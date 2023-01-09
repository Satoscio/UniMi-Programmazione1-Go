/*
	Scrivere un programma primo.go che, dato un numero intero su standard input,
	determina se il numero è primo.

	Suggerimento: occorre determinare il primo numero che è un divisore (se c'è).
	Domanda: dato in input n, quante iterazioni dovrò fare al massimo?
*/

package main

import "fmt"

func main() {
	var n, divs int
	fmt.Scan(&n)

	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			divs++
		}
	}
	if divs == 0 {
		fmt.Println("Numero primo")
	} else {
		fmt.Println("Numero non primo")
	}
}
