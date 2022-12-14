/*
	Scrivere un programma num2text.go per convertire un numero intero
	non negativo nella sequenza delle parole corrispondenti alle sue cifre.

	Il programma legge un intero non negativo da standard input, per
	ogni nuova (non incontrata finora) cifra del numero chiede il nome
	corrispondente (e alimenta un dizionario), e infine stampa la
	sequenzadelle parole corrispondenti alle sue cifre.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var num int
	var numS string
	diz := map[string]string{}

	fmt.Print("Numero: ")
	fmt.Scan(&num)
	nms := strconv.Itoa(num)
	nsl := strings.Split(nms, "")

	for _, n := range nsl {
		if _, ok := diz[n]; !ok {
			fmt.Printf("Parola per %s: ", n)
			fmt.Scan(&numS)
			diz[n] = numS 
		}
	}
	
	for _, n := range nsl {
		fmt.Printf("%s ", diz[n])
	}
	fmt.Println()
}
