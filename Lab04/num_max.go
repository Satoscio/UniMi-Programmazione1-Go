/*
	Scrivere un programma num_max.go che legge una sequenza di
	10 interi positivi e stampa il massimo intero letto e quante
	volte tale massimo compare nella sequenza.

	E` possibile risolvere il problema senza tenere in memoria
	la sequenza di interi. Che tipo di composizione occorrerà
	per mettere insieme il calcolo del massimo e il conteggio
	delle occorrenze di tale massimo?
*/

package main

import "fmt"

func main() {
	const NUMERI = 10
	var n, min, max, occ int
	for i := 0; i < NUMERI; i++ {
		fmt.Scan(&n)
		if n > min {
			max, min = n, n
			occ = 0
		}
		if n == max {
			occ++
		}
	}
	fmt.Printf("Il max è %d e appare %d volte\n", max, occ)
}
