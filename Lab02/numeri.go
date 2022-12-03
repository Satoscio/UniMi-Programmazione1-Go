/*
    Specifiche: Scrivere un programma numeri.go che riceve in ingresso un numero
    intero e stampa "divisibile per 10" se il numero è divisibile per 10,
    "positivo" se il numero è positivo, "positivo divisibile per 10" se è sia
    divisibile per 10 che positivo, niente altrimenti.
*/

package main

import "fmt"

func main() {
    var n int
    fmt.Print("numero? ")
	fmt.Scan(&n)

    if n >= 0 {
		fmt.Print("positivo ")
	}
	if n % 10 == 0 {
		fmt.Print("divisibile per 10")
	}

    fmt.Println()
}
