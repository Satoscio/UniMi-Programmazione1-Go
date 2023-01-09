/*
	Scrivere un programma cesare.go che legge da standard input un valore intero
	non negativo k (la chiave di cifratura) e una sequenza di lettere minuscole
	consecutive (sulla stessa riga e senza spazi) terminate da invio.

	Il programma stampa la sequenza letta cifrata secondo il cifrario di Cesare,
	usando come chiave k (quella fornita dall'utente):

	ogni lettera del testo in chiaro è sostituita nel testo cifrato dalla lettera
	che si trova k posizioni dopo nell'alfabeto, ritornando alla lettera a dopo la zeta.
*/

package main

import "fmt"

func main() {
	var s string
	var r rune
	var n int
	fmt.Print("Chiave: ")
	fmt.Scan(&n)
	fmt.Print("Caratteri da cifrare: ")
	for {
		fmt.Scanf("%c", &r)
		if r == 13 || r == 10 {
			break
		}
		if n > 26 {
			n %= 26
		}
		if r+rune(n) > 'z' {
			t := (r + rune(n)) - 'z'
			r = 'a' + t - 1
			//fmt.Println("qui")
		} else {
			r += rune(n)
		}
		s += string(r)
	}
	fmt.Println("Testo cifrato:", s)
}
