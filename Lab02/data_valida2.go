/*
	Scrivere una seconda versione data_valida2.go che
	legge due interi gg e mm rappresentanti giorno e
	mese dell’anno, e verifica che sia una data valida.

	In questa seconda versione si tenga conto del fatto
	che solo i mesi 1, 3, 5, 7, 8, 10, 12 hanno 31 giorni,
	che i mesi 4, 6, 9, 11 ne hanno 30, e si assuma che
	febbraio ne abbia sempre 28.

	Il programma stampa "data valida" / "data non valida".
*/

package main

import (
	"fmt"
)

func main() {
	var gg, mm int
	fmt.Scan(&gg, &mm)


	switch mm {
		case 1, 3, 5, 7, 8, 10, 12:
			if gg < 1 || gg > 31 {
				fmt.Println("Data non valida")
			} else {
				fmt.Println("Data valida")
			}
		case 4, 6, 9, 11:
			if gg < 1 || gg > 30 {
				fmt.Println("Data non valida")
			} else {
				fmt.Println("Data valida")
			}
		case 2:
			if gg < 1 || gg > 28 {
				fmt.Println("Data non valida")
			} else {
				fmt.Println("Data valida")
			}
		default:
			fmt.Println("Data non valida")
	}
}