/*
	Scrivere un programma bin2ten.go che converte un intero composto di soli 0 e 1
	nel valore corrispondente al numero binario rappresentato (es. 101 -> 5).
	Nel caso il numero contenga altre cifre, il programma stampa un messaggio di errore.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var bin, dec int
	fmt.Scan(&bin)

	for i := 0; i < 16; i++ {
		n := bin % 10
		bin /= 10

		if n != 0 && n != 1 {
			fmt.Println("Non è un numero binario!")
			break
		}

		dec += n * int(math.Pow(float64(2), float64(i)))
	}
	fmt.Println(dec)
}
