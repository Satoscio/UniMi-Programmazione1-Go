/*
	Scrivere un programma data_valida1.go che legge due
	interi gg e mm rappresentanti giorno e mese dell’anno,
	e verifica che sia una data valida.
	In questa versione si assuma che tutti i mesi abbiano 31 giorni.
	Il programma stampa "data valida" / "data non valida".
*/

package main

import "fmt"

func main() {
	var gg, mm int
	fmt.Scan(&gg, &mm)
	if gg < 1 || gg > 31 {
		fmt.Println("Data non valida")
	} else {
		if mm < 1 || mm > 12 {
			fmt.Println("Data non valida")
		} else {
			fmt.Println("Data valida")
		}
	}

}