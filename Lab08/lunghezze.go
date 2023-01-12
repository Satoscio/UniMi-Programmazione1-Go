/*
   Scrivere un programma lunghezze.go che legge riga per riga un testo da
   standard input (potete usare la ridirezione), terminato da EOF, e stampa
   quante parole ci sono di lunghezza 1, quante di lunghezza 2, ecc.

   Il programma è dotato di una funzione
   - func aggiornaConteggio(m map[int]int, riga string)
   che aggiorna la mappa dei conteggi e che deve essere usata dal main.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func aggiornaConteggio(m map[int]int, riga string) {
	parole := strings.Split(riga, " ")
	for _, p := range parole {
		m[len(p)]++
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lungh := map[int]int{}
	numeri := []int{}

	for scanner.Scan() {
		aggiornaConteggio(lungh, scanner.Text())
	}
	delete(lungh, 0)
	for i := range lungh {
		numeri = append(numeri, i)
	}
	sort.Ints(numeri)
	
	for _, i := range numeri {
		fmt.Printf("Lunghezza: %d, Parole: %d\n", i, lungh[i])
	}
}
