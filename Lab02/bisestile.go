/*
	Scrivere un programma bisestile.go che legge un intero n
	corrispondente all’anno di una data, e determina se l’anno
	è bisestile o no (stampa "bisestile" o "non bisestile").
*/

package main

import "fmt"

func main() {
	var anno int
	fmt.Scan(&anno)
	if anno % 4 == 0 {
		if anno % 100 == 0 {
			if anno % 400 == 0 {
				fmt.Println("Anno bisestile secolare.")
			} else {
				fmt.Println("Anno non bisestile secolare.")
			}
		} else {
			fmt.Println("Anno bisestile")
		}
	} else {
		fmt.Println("Anno non bisestile")
	}
}