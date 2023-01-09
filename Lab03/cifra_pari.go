/*
	Scrivere un programma cifra_pari.go che, dato un intero da standard input,
	determina e stampa in che posizione (procedendo da destra a sinistra) si trova
	la prima cifra pari del numero. Se il numero non contiene cifre pari, il programma stampa -1.
*/

package main

import (
	"fmt"
)

func main() {
	var n, i int
	fmt.Scan(&n)
	for {
		x := n % 10
		n /= 10

		if x%2 == 0 {
			fmt.Println("Pos:", i)
			break
		}

		if n == 0 {
			fmt.Println("Non contiene pari")
			break
		}
		i++
	}
}
