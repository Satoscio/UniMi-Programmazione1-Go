/*
	Scrivere un programma capitale.go che legge da stardard input tre valori float64
	che rappresentano un capitale iniziale (es. 1000.50 euro), un tasso di interesse
	annuale (es. 1.3%) e un obiettivo (es. 2000 euro), e calcola quanti anni occorrono
	per arrivare a (o superare) l'obiettivo.
*/

package main

import "fmt"

func main() {
	var cap_in, inter, ob float64
	var anni int
	fmt.Scan(&cap_in, &inter, &ob)
	for {
		cap_in += (cap_in / 100 * inter)
		anni++
		if cap_in >= ob {
			break
		}
	}
	fmt.Printf("Importo finale: €%.2f raggiunto dopo %d anni\n", cap_in, anni)
}
