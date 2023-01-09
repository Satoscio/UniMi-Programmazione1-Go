/*

	Scrivere un programma andamento.go che legge da
	tastiera una serie (di almeno un elemento) di numeri
	interi > -1 e stampa "+" ogni volta che il nuovo
	valore è maggiore o uguale al precedente e "-"
	altrimenti.
	Si ferma quando il numero in input è -1 e stampa
	la somma di tutti i numeri letti.

	Esempio
	$ go run andamento.go
	2 4 7 3 9 1 5 -1
	++-+-+
	somma: 31

*/

package main

import "fmt"

func main() {
	var prec, n, somma int
	for {
		prec = n
		fmt.Scan(&n)

		if n == -1 {
			break
		}
		somma += n
		if prec <= n && prec != 0 {
			fmt.Print("+")
		} else if prec == 0 {
			fmt.Print()
		} else {
			fmt.Print("-")
		}

	}
	fmt.Printf("\nSomma: %d\n", somma)
}
