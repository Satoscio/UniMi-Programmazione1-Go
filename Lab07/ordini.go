/*
	Scrivere un programma ordini.go che legge da standard input una serie di
	stringhe che descrivono ordini nel formato prezzo#quantità#sconto e stampa
	il totale finale da pagare.

	Prezzo, quantità e sconto sono float; prezzo indica il prezzo unitario del
	prodotto, quantità indica la quantità acquistata e sconto è lo sconto
	applicato per quel prodotto, espresso come float tra 0 e 1
	(ad esempio 0.2 indica uno sconto del 20%).
	Il programma termina la lettura quando incontra EOF (ctrl D su riga nuova).
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var totale float64
	ordini := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ordini = append(ordini, scanner.Text())
	}
	
	for _, o := range ordini {
		ord1 := strings.Split(o, "#")
		pr, _ := strconv.ParseFloat(ord1[0], 64)
		qa, _ := strconv.ParseFloat(ord1[1], 64)
		sc, _ := strconv.ParseFloat(ord1[2], 64)
		totale += (pr * qa) - ((pr * qa) * sc) 
	}
	
	fmt.Println("Il totale da pagare è", totale)
}