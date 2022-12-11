/*
	Scrivere un programma extractions.go con due funzioni:

	estraiPari(in []int) (out []int) che prende una slice di interi e ne
	restituisce una che contiene solo i numeri di quella in ingresso che
	sono numeri pari.

	rimuoviMultipli(m int, in []int) (out []int) che prende un intero e
	una slice di interi e ne restituisce una senza i multipli del numero
	passato di quella in ingresso. Es.: rimuoviMultipli(5, in) restituisce,
	a partire da in, una slice senza i multipli di 5.
*/

package main

import "fmt"

func estraiPari(in []int) (out []int) {
	for _, n := range in {
		if n%2 == 0 {
			out = append(out, n)
		}
	}
	return
}

func rimuoviMultipli(m int, in []int) (out []int) {
	for _, n := range in {
		if n%m != 0 {
			out = append(out, n)
		}
	}
	return
}

func main() {
	var m int
	interi := []int{2, 15, 23, 4, 7, 9, 375, 888, 902, 18}
	fmt.Printf("Slice:\t\t\t%v\n", interi)
	fmt.Printf("Pari estratti:\t\t%v\n", estraiPari(interi))
	fmt.Print("Togli multipli di: ")
	fmt.Scan(&m)
	fmt.Printf("Senza multipli di %d:\t%v", m, rimuoviMultipli(m, interi))
}
