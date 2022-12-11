/*
	Scrivere un programma stampaAlternata.go che legge da standard input del testo su
	più righe (terminato da EOF) e stampa prima le righe pari e poi le righe dispari
	(considerate la prima riga del testo la riga 1 (e non 0)).
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	testi := []string{}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		testi = append(testi, scanner.Text())
	}

	fmt.Print("\nRighe pari\n")
	for i := 0; i < len(testi); i += 2 {
		fmt.Println(testi[i])
	}
	fmt.Print("\nRighe dispari\n")
	for i := 1; i < len(testi); i += 2 {
		fmt.Println(testi[i])
	}

}
